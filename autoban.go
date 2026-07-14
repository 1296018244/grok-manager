package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Ported from https://github.com/akihitohyh/xai-autoban (MIT) + local enhancements:
// usage.handle → isolate bad xAI creds; scheduler.pick → skip them; optional bans.json persist.

const (
	bansFileName = "bans.json"

	banUnauthorizedLong  = 24 * time.Hour
	banUnauthorizedShort = 2 * time.Hour // 401 with vault SSO — leave room for auto-refresh
	banPayment           = 7 * 24 * time.Hour
	banForbidden         = 24 * time.Hour
	// 429: fixed 2h isolation (shared pools; don't guess window start).
	banRateFixed = 2 * time.Hour
)

type banEntry struct {
	StatusCode int       `json:"status_code"`
	Reason     string    `json:"reason"`
	Source     string    `json:"source,omitempty"` // usage | scan | import
	Email      string    `json:"email,omitempty"`
	Label      string    `json:"label,omitempty"`
	BannedAt   time.Time `json:"banned_at"`
	ResetAt    time.Time `json:"reset_at"`
	FailCount  int       `json:"fail_count,omitempty"`
}

type banState struct {
	mu         sync.Mutex
	bans       map[string]banEntry
	emailIndex map[string]map[string]struct{} // email(lower) → authIDs
	path       string
	dirty      bool
	persist    bool // default true
}

var runtimeBans = &banState{
	bans:       map[string]banEntry{},
	emailIndex: map[string]map[string]struct{}{},
	persist:    true,
}

func (s *banState) set(authID string, entry banEntry) {
	authID = strings.TrimSpace(authID)
	if authID == "" {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.bans == nil {
		s.bans = make(map[string]banEntry)
	}
	if current, ok := s.bans[authID]; ok {
		entry.FailCount = current.FailCount + 1
		// recheck429 must always apply the new window (+2h from probe).
		if entry.Source != "recheck429" && current.ResetAt.After(entry.ResetAt) {
			// Keep longer isolation window, but refresh metadata.
			if entry.Email != "" {
				current.Email = entry.Email
			}
			if entry.Label != "" {
				current.Label = entry.Label
			}
			if entry.StatusCode != 0 {
				current.StatusCode = entry.StatusCode
				current.Reason = entry.Reason
			}
			current.FailCount = entry.FailCount
			current.Source = firstNonEmpty(entry.Source, current.Source)
			s.bans[authID] = current
			s.indexEmailLocked(authID, current.Email)
			s.dirty = true
			go saveBansAsync()
			return
		}
		if entry.Email == "" {
			entry.Email = current.Email
		}
		if entry.Label == "" {
			entry.Label = current.Label
		}
	} else if entry.FailCount == 0 {
		entry.FailCount = 1
	}
	s.bans[authID] = entry
	s.indexEmailLocked(authID, entry.Email)
	s.dirty = true
	go saveBansAsync()
}

func (s *banState) indexEmailLocked(authID, email string) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" {
		return
	}
	if s.emailIndex == nil {
		s.emailIndex = map[string]map[string]struct{}{}
	}
	if s.emailIndex[email] == nil {
		s.emailIndex[email] = map[string]struct{}{}
	}
	s.emailIndex[email][authID] = struct{}{}
}

func (s *banState) unindexLocked(authID string, email string) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" || s.emailIndex == nil {
		return
	}
	if m := s.emailIndex[email]; m != nil {
		delete(m, authID)
		if len(m) == 0 {
			delete(s.emailIndex, email)
		}
	}
}

func (s *banState) active(authID string, now time.Time) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	entry, ok := s.bans[authID]
	if !ok {
		return false
	}
	if now.Before(entry.ResetAt) {
		return true
	}
	// 429: sticky after expiry until auto-recheck decides (still 429 → +2h, else unban).
	// Other codes: drop from isolation when window ends.
	if entry.StatusCode == http.StatusTooManyRequests {
		return true
	}
	s.unindexLocked(authID, entry.Email)
	delete(s.bans, authID)
	s.dirty = true
	go saveBansAsync()
	return false
}

func (s *banState) clear(authID string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	authID = strings.TrimSpace(authID)
	entry, ok := s.bans[authID]
	if !ok {
		return false
	}
	s.unindexLocked(authID, entry.Email)
	delete(s.bans, authID)
	s.dirty = true
	go saveBansAsync()
	return true
}

func (s *banState) clearAll() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := len(s.bans)
	s.bans = make(map[string]banEntry)
	s.emailIndex = map[string]map[string]struct{}{}
	s.dirty = true
	go saveBansAsync()
	return n
}

func (s *banState) clearStatus(status int) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	removed := 0
	for id, entry := range s.bans {
		if entry.StatusCode == status {
			s.unindexLocked(id, entry.Email)
			delete(s.bans, id)
			removed++
		}
	}
	if removed > 0 {
		s.dirty = true
		go saveBansAsync()
	}
	return removed
}

func (s *banState) clearMany(authIDs []string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	removed := 0
	for _, id := range authIDs {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		if entry, ok := s.bans[id]; ok {
			s.unindexLocked(id, entry.Email)
			delete(s.bans, id)
			removed++
		}
	}
	if removed > 0 {
		s.dirty = true
		go saveBansAsync()
	}
	return removed
}

func (s *banState) clearByEmail(email string) int {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" {
		return 0
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	removed := 0
	// via index
	if ids := s.emailIndex[email]; len(ids) > 0 {
		for id := range ids {
			if entry, ok := s.bans[id]; ok {
				s.unindexLocked(id, entry.Email)
				delete(s.bans, id)
				removed++
			}
		}
	}
	// sweep in case index missed
	for id, entry := range s.bans {
		if strings.ToLower(entry.Email) == email {
			s.unindexLocked(id, entry.Email)
			delete(s.bans, id)
			removed++
		}
	}
	if removed > 0 {
		s.dirty = true
		go saveBansAsync()
	}
	return removed
}

func (s *banState) snapshot(now time.Time) map[string]banEntry {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make(map[string]banEntry, len(s.bans))
	changed := false
	for id, entry := range s.bans {
		if now.Before(entry.ResetAt) {
			out[id] = entry
			continue
		}
		// Keep expired 429 until recheck; purge other expired bans.
		if entry.StatusCode == http.StatusTooManyRequests {
			out[id] = entry
			continue
		}
		s.unindexLocked(id, entry.Email)
		delete(s.bans, id)
		changed = true
	}
	if changed {
		s.dirty = true
		go saveBansAsync()
	}
	return out
}

// hasExpired429 reports whether any sticky-expired 429 is waiting for recheck.
func (s *banState) hasExpired429(now time.Time) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, e := range s.bans {
		if e.StatusCode == http.StatusTooManyRequests && !e.ResetAt.After(now) {
			return true
		}
	}
	return false
}

func (s *banState) count() int {
	return len(s.snapshot(time.Now()))
}

// failCountForStatus returns max FailCount among active bans matching email+status (or 0).
func (s *banState) failCountForStatus(email string, status int) int {
	email = strings.ToLower(strings.TrimSpace(email))
	s.mu.Lock()
	defer s.mu.Unlock()
	max := 0
	now := time.Now()
	for _, e := range s.bans {
		if e.StatusCode != status {
			continue
		}
		// Treat sticky-expired 429 as still active for fail-count stats.
		if !e.ResetAt.After(now) && e.StatusCode != http.StatusTooManyRequests {
			continue
		}
		if email != "" && strings.ToLower(e.Email) != email {
			continue
		}
		if email == "" {
			continue
		}
		if e.FailCount > max {
			max = e.FailCount
		}
	}
	return max
}

// ---- persistence ----

type bansFile struct {
	SavedAt string              `json:"saved_at"`
	Bans    map[string]banEntry `json:"bans"`
}

func bansFilePath() string {
	runtimeBans.mu.Lock()
	defer runtimeBans.mu.Unlock()
	if runtimeBans.path != "" {
		return runtimeBans.path
	}
	return resolvePluginDataPath(bansFileName, &runtimeBans.path)
}

func loadBansOnStart() {
	for _, p := range pluginDataCandidates(bansFileName) {
		raw, err := os.ReadFile(p)
		if err != nil || len(strings.TrimSpace(string(raw))) == 0 {
			continue
		}
		var f bansFile
		if err := json.Unmarshal(raw, &f); err != nil {
			continue
		}
		now := time.Now()
		runtimeBans.mu.Lock()
		runtimeBans.path = p
		runtimeBans.bans = map[string]banEntry{}
		runtimeBans.emailIndex = map[string]map[string]struct{}{}
		for id, e := range f.Bans {
			if e.ResetAt.IsZero() || !e.ResetAt.After(now) {
				continue
			}
			runtimeBans.bans[id] = e
			runtimeBans.indexEmailLocked(id, e.Email)
		}
		runtimeBans.dirty = false
		runtimeBans.mu.Unlock()
		return
	}
}

func saveBansAsync() {
	// debounce-ish: small sleep then flush
	time.Sleep(80 * time.Millisecond)
	saveBans()
}

func saveBans() {
	runtimeBans.mu.Lock()
	if !runtimeBans.persist || !runtimeBans.dirty {
		runtimeBans.mu.Unlock()
		return
	}
	path := runtimeBans.path
	if path == "" {
		path = resolvePluginDataPath(bansFileName, &runtimeBans.path)
	}
	// copy under lock
	snap := make(map[string]banEntry, len(runtimeBans.bans))
	now := time.Now()
	for id, e := range runtimeBans.bans {
		if e.ResetAt.After(now) {
			snap[id] = e
		}
	}
	runtimeBans.dirty = false
	runtimeBans.mu.Unlock()

	payload := bansFile{SavedAt: time.Now().UTC().Format(time.RFC3339), Bans: snap}
	raw, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return
	}
	_ = os.MkdirAll(filepath.Dir(path), 0o755)
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, raw, 0o644); err != nil {
		return
	}
	if err := os.Rename(tmp, path); err != nil {
		_ = os.WriteFile(path, raw, 0o644)
		_ = os.Remove(tmp)
	}
}

// ---- CPA hooks ----

type usageRecord struct {
	Provider        string `json:"Provider"`
	ProviderAlt     string `json:"provider"`
	AuthID          string `json:"AuthID"`
	AuthIDAlt       string `json:"auth_id"`
	AuthIndex       string `json:"AuthIndex"`
	Failed          bool   `json:"Failed"`
	FailedAlt       *bool  `json:"failed"`
	Failure         struct {
		StatusCode int    `json:"StatusCode"`
		StatusAlt  int    `json:"status_code"`
		Body       string `json:"Body"`
	} `json:"Failure"`
	FailureAlt struct {
		StatusCode int `json:"status_code"`
	} `json:"failure"`
	ResponseHeaders http.Header `json:"ResponseHeaders"`
	HeadersAlt      http.Header `json:"response_headers"`
	Email           string      `json:"Email"`
	EmailAlt        string      `json:"email"`
	Alias           string      `json:"Alias"`
	Source          string      `json:"Source"`
}

func (r usageRecord) provider() string  { return firstNonEmpty(r.Provider, r.ProviderAlt) }
func (r usageRecord) authID() string    { return firstNonEmpty(r.AuthID, r.AuthIDAlt, r.AuthIndex) }
func (r usageRecord) failed() bool {
	if r.FailedAlt != nil {
		return *r.FailedAlt
	}
	return r.Failed
}
func (r usageRecord) statusCode() int {
	if r.Failure.StatusCode != 0 {
		return r.Failure.StatusCode
	}
	if r.Failure.StatusAlt != 0 {
		return r.Failure.StatusAlt
	}
	return r.FailureAlt.StatusCode
}
func (r usageRecord) headers() http.Header {
	if len(r.ResponseHeaders) > 0 {
		return r.ResponseHeaders
	}
	return r.HeadersAlt
}
func (r usageRecord) email() string { return firstNonEmpty(r.Email, r.EmailAlt) }

type schedulerPickRequest struct {
	Provider      string                   `json:"Provider"`
	Candidates    []schedulerAuthCandidate `json:"Candidates"`
	ProviderAlt   string                   `json:"provider"`
	CandidatesAlt []schedulerAuthCandidate `json:"candidates"`
}

type schedulerAuthCandidate struct {
	ID          string `json:"ID"`
	IDAlt       string `json:"id"`
	Provider    string `json:"Provider"`
	ProviderAlt string `json:"provider"`
	Priority    int    `json:"Priority"`
	PriorityAlt int    `json:"priority"`
	Email       string `json:"Email"`
	EmailAlt    string `json:"email"`
}

func (c schedulerAuthCandidate) id() string       { return firstNonEmpty(c.ID, c.IDAlt) }
func (c schedulerAuthCandidate) provider() string { return firstNonEmpty(c.Provider, c.ProviderAlt) }
func (c schedulerAuthCandidate) priority() int {
	if c.Priority != 0 {
		return c.Priority
	}
	return c.PriorityAlt
}
func (c schedulerAuthCandidate) email() string { return firstNonEmpty(c.Email, c.EmailAlt) }

type schedulerPickResponse struct {
	AuthID  string `json:"AuthID"`
	Handled bool   `json:"Handled"`
}

func handleUsage(raw []byte) ([]byte, error) {
	if len(raw) == 0 {
		return okEnvelope(map[string]any{})
	}
	var record usageRecord
	if err := json.Unmarshal(raw, &record); err != nil {
		return okEnvelope(map[string]any{})
	}
	if !strings.EqualFold(record.provider(), xaiProvider) || !record.failed() {
		return okEnvelope(map[string]any{})
	}
	authID := record.authID()
	if authID == "" {
		return okEnvelope(map[string]any{})
	}
	email := record.email()
	now := time.Now()
	entry, ok := classifyFailure(record.statusCode(), record.headers(), now, email)
	if !ok {
		return okEnvelope(map[string]any{})
	}
	entry.Source = "usage"
	entry.Email = email
	entry.Label = firstNonEmpty(record.Alias, record.Source)
	// Ban under canonical AuthID + AuthIndex aliases.
	runtimeBans.set(authID, entry)
	if idx := strings.TrimSpace(record.AuthIndex); idx != "" && idx != authID {
		runtimeBans.set(idx, entry)
	}
	return okEnvelope(map[string]any{})
}

func classifyFailure(status int, headers http.Header, now time.Time, email string) (banEntry, bool) {
	entry := banEntry{StatusCode: status, BannedAt: now, Email: email}
	switch status {
	case http.StatusUnauthorized:
		entry.Reason = "unauthorized"
		// Short ban when vault can auto-refresh; long when no SSO.
		if email != "" {
			if _, ok := vaultLookupByEmail(email); ok {
				entry.Reason = "unauthorized_vault"
				entry.ResetAt = now.Add(banUnauthorizedShort)
				break
			}
		}
		entry.ResetAt = now.Add(banUnauthorizedLong)
	case http.StatusPaymentRequired:
		entry.Reason = "payment_required"
		entry.ResetAt = now.Add(banPayment)
	case http.StatusForbidden:
		entry.Reason = "forbidden"
		entry.ResetAt = now.Add(banForbidden)
	case http.StatusTooManyRequests:
		entry.Reason = "rate_limited_2h"
		entry.ResetAt = now.Add(banRateFixed)
	default:
		return banEntry{}, false
	}
	return entry, true
}

func rateLimitReset(headers http.Header, now time.Time) time.Time {
	if headers == nil {
		return time.Time{}
	}
	if raw := strings.TrimSpace(headers.Get("Retry-After")); raw != "" {
		if seconds, err := strconv.ParseInt(raw, 10, 64); err == nil && seconds > 0 {
			return now.Add(time.Duration(seconds) * time.Second)
		}
		if parsed, err := http.ParseTime(raw); err == nil && parsed.After(now) {
			return parsed
		}
	}
	for _, key := range []string{"x-ratelimit-reset", "x-ratelimit-reset-requests", "X-RateLimit-Reset"} {
		raw := strings.TrimSpace(headers.Get(key))
		value, err := strconv.ParseInt(raw, 10, 64)
		if err != nil || value <= 0 {
			continue
		}
		if value > 1_000_000_000_000 {
			value /= 1000
		}
		reset := time.Unix(value, 0)
		if reset.After(now) {
			return reset
		}
	}
	return time.Time{}
}

func handleSchedulerPick(raw []byte) ([]byte, error) {
	var req schedulerPickRequest
	if err := json.Unmarshal(raw, &req); err != nil {
		return nil, err
	}
	cands := req.Candidates
	if len(cands) == 0 {
		cands = req.CandidatesAlt
	}
	now := time.Now()
	available := make([]schedulerAuthCandidate, 0, len(cands))
	for _, c := range cands {
		if strings.EqualFold(c.provider(), xaiProvider) && isBannedCandidate(c, now) {
			continue
		}
		available = append(available, c)
	}
	if len(available) == len(cands) || len(available) == 0 {
		return okEnvelope(schedulerPickResponse{Handled: false})
	}
	chosen := available[0]
	for _, c := range available[1:] {
		if c.priority() > chosen.priority() {
			chosen = c
		}
	}
	return okEnvelope(schedulerPickResponse{AuthID: chosen.id(), Handled: true})
}

func isBannedCandidate(c schedulerAuthCandidate, now time.Time) bool {
	if runtimeBans.active(c.id(), now) {
		return true
	}
	// Also match by email aliases recorded from scan/usage.
	em := strings.ToLower(strings.TrimSpace(c.email()))
	if em == "" {
		return false
	}
	runtimeBans.mu.Lock()
	ids := runtimeBans.emailIndex[em]
	runtimeBans.mu.Unlock()
	for id := range ids {
		if runtimeBans.active(id, now) {
			return true
		}
	}
	return false
}

// noteScanBan isolates after active probe. Prefer host AuthID / ID / AuthIndex.
func noteScanBan(res probeResult) {
	entry, ok := classifyFailure(res.HTTPStatus, nil, time.Now(), res.Email)
	if !ok {
		return
	}
	entry.Source = "scan"
	entry.Email = res.Email
	entry.Label = firstNonEmpty(res.Name, res.File)
	ids := []string{res.AuthID, res.AuthIndex, res.Name, res.File}
	if res.Path != "" {
		ids = append(ids, res.Path, filepathBase(res.Path))
	}
	seen := map[string]struct{}{}
	for _, id := range ids {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		if _, dup := seen[id]; dup {
			continue
		}
		seen[id] = struct{}{}
		runtimeBans.set(id, entry)
	}
}

func filepathBase(p string) string {
	p = strings.ReplaceAll(p, "\\", "/")
	if i := strings.LastIndex(p, "/"); i >= 0 {
		return p[i+1:]
	}
	return p
}

// noteSSOSuccess is stubbed in stubs.go for public build (no SSO convert).

// ---- status / management ----

type banInfo struct {
	AuthID           string `json:"auth_id"`
	StatusCode       int    `json:"status_code"`
	Reason           string `json:"reason"`
	Source           string `json:"source,omitempty"`
	Email            string `json:"email,omitempty"`
	Label            string `json:"label,omitempty"`
	BannedAt         string `json:"banned_at"`
	ResetAt          string `json:"reset_at"`
	RemainingSeconds int64  `json:"remaining_seconds"`
	FailCount        int    `json:"fail_count,omitempty"`
}

type autobanStatus struct {
	Plugin     string           `json:"plugin"`
	Version    string           `json:"version"`
	Count      int              `json:"count"`
	Match      int              `json:"match"`
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	Pages      int              `json:"pages"`
	Filter     string           `json:"filter,omitempty"`
	Q          string           `json:"q,omitempty"`
	ByCode     map[int]int      `json:"by_code"`
	Bans       []banInfo        `json:"bans"`
	Note       string           `json:"note,omitempty"`
	BansPath   string           `json:"bans_path,omitempty"`
	Persisted  bool             `json:"persisted"`
	Recheck429 recheck429Result `json:"recheck_429,omitempty"`
	Due429     int              `json:"due_429"` // expired sticky 429 waiting recheck
}

func autobanSnapshot(q url.Values) autobanStatus {
	pq := parsePageQuery(q)
	// Allow status via filter too (e.g. filter=429).
	if pq.Status == 0 && pq.Filter != "" && pq.Filter != "all" {
		if n, err := strconv.Atoi(pq.Filter); err == nil {
			pq.Status = n
		}
	}
	now := time.Now()
	snap := runtimeBans.snapshot(now)
	items := make([]banInfo, 0, len(snap))
	byCode := map[int]int{}
	due429 := 0
	for id, entry := range snap {
		info := banInfo{
			AuthID:     id,
			StatusCode: entry.StatusCode,
			Reason:     entry.Reason,
			Source:     entry.Source,
			Email:      entry.Email,
			Label:      entry.Label,
			BannedAt:   entry.BannedAt.Format(time.RFC3339),
			ResetAt:    entry.ResetAt.Format(time.RFC3339),
			RemainingSeconds: func() int64 {
				sec := int64(entry.ResetAt.Sub(now).Seconds())
				if sec < 0 {
					return 0
				}
				return sec
			}(),
			FailCount: entry.FailCount,
		}
		byCode[entry.StatusCode]++
		if entry.StatusCode == http.StatusTooManyRequests && !entry.ResetAt.After(now) {
			due429++
		}
		if pq.Status > 0 && entry.StatusCode != pq.Status {
			continue
		}
		if pq.Q != "" {
			qq := strings.ToLower(pq.Q)
			blob := strings.ToLower(strings.Join([]string{
				id, entry.Email, entry.Reason, entry.Label, entry.Source,
			}, " "))
			if !strings.Contains(blob, qq) {
				continue
			}
		}
		items = append(items, info)
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].ResetAt == items[j].ResetAt {
			return items[i].AuthID < items[j].AuthID
		}
		return items[i].ResetAt < items[j].ResetAt
	})
	pageItems, match, pages, page := slicePage(items, pq.Page, pq.PageSize)
	filterLabel := "all"
	if pq.Status > 0 {
		filterLabel = strconv.Itoa(pq.Status)
	}
	return autobanStatus{
		Plugin:     pluginName,
		Version:    pluginVersion,
		Count:      len(snap),
		Match:      match,
		Page:       page,
		PageSize:   pq.PageSize,
		Pages:      pages,
		Filter:     filterLabel,
		Q:          pq.Q,
		ByCode:     byCode,
		Bans:       pageItems,
		BansPath:   bansFilePath(),
		Persisted:  true,
		Recheck429: recheck429Public(),
		Due429:     due429,
	}
}

// ---- 429 recheck (manual all | auto on expiry) ----
//
// Policy: 429 is isolated for 2h. When the window ends the ban stays sticky
// (scheduler still skips it) until a probe runs:
//   still 429 → refresh +2h
//   401/402/403 → reclassify
//   other OK  → unban
// Manual button still probes every active 429.

const recheck429Poll = 30 * time.Second

type recheck429Item struct {
	AuthID     string `json:"auth_id"`
	Email      string `json:"email,omitempty"`
	HTTPStatus int    `json:"http_status"`
	Action     string `json:"action"` // still_429 | unbanned | reclassified | skipped | error
	Detail     string `json:"detail,omitempty"`
}

type recheck429Result struct {
	Running      bool             `json:"running"`
	Manual       bool             `json:"manual"`
	Mode         string           `json:"mode,omitempty"` // manual | expiry
	Checked      int              `json:"checked"`
	Still429     int              `json:"still_429"`
	Unbanned     int              `json:"unbanned"`
	Reclassified int              `json:"reclassified"`
	Skipped      int              `json:"skipped"`
	Errors       int              `json:"errors"`
	Message      string           `json:"message"`
	LastRun      string           `json:"last_run,omitempty"`
	NextHourly   string           `json:"next_hourly,omitempty"` // legacy field: next auto poll hint
	Details      []recheck429Item `json:"details,omitempty"`
	Status       *autobanStatus   `json:"status,omitempty"`
}

var (
	recheck429Mu      sync.Mutex
	recheck429Running bool
	recheck429Last    recheck429Result
	recheck429Once    sync.Once
)

func recheck429Public() recheck429Result {
	recheck429Mu.Lock()
	defer recheck429Mu.Unlock()
	out := recheck429Last
	out.Running = recheck429Running
	// Don't dump details on status poll (can be large).
	out.Details = nil
	out.Status = nil
	return out
}

func startRecheck429Loop() {
	recheck429Once.Do(func() {
		go func() {
			for {
				time.Sleep(recheck429Poll)
				recheck429Mu.Lock()
				busy := recheck429Running
				recheck429Mu.Unlock()
				if busy {
					continue
				}
				// Only when some 429 isolation window has ended.
				if !runtimeBans.hasExpired429(time.Now()) {
					continue
				}
				_, _ = runRecheck429(false)
			}
		}()
	})
}

func handleRecheck429() ([]byte, error) {
	res, err := runRecheck429(true)
	if err != nil {
		if err.Error() == "busy" {
			return jsonErrorEnvelope(http.StatusConflict, "busy", "429 recheck already running")
		}
		return jsonErrorEnvelope(http.StatusInternalServerError, "recheck_failed", err.Error())
	}
	return jsonManagementEnvelope(http.StatusOK, res)
}

func runRecheck429(manual bool) (recheck429Result, error) {
	recheck429Mu.Lock()
	if recheck429Running {
		recheck429Mu.Unlock()
		return recheck429Result{}, fmt.Errorf("busy")
	}
	recheck429Running = true
	recheck429Mu.Unlock()

	defer func() {
		recheck429Mu.Lock()
		recheck429Running = false
		recheck429Mu.Unlock()
	}()

	mode := "expiry"
	if manual {
		mode = "manual"
	}
	res := recheck429Result{
		Running: true,
		Manual:  manual,
		Mode:    mode,
		Message: "probing 429 bans (" + mode + ")",
	}

	now := time.Now()
	snap := runtimeBans.snapshot(now)
	var targets []struct {
		id    string
		entry banEntry
	}
	for id, entry := range snap {
		if entry.StatusCode != http.StatusTooManyRequests {
			continue
		}
		// Auto loop: only expired (sticky) 429 — verify whether isolation still needed.
		// Manual: every 429 currently isolated.
		if !manual && entry.ResetAt.After(now) {
			continue
		}
		targets = append(targets, struct {
			id    string
			entry banEntry
		}{id, entry})
	}
	sort.Slice(targets, func(i, j int) bool { return targets[i].id < targets[j].id })
	res.Checked = len(targets)
	if len(targets) == 0 {
		res.Running = false
		if manual {
			res.Message = "no 429 bans"
		} else {
			res.Message = "no expired 429"
		}
		res.LastRun = now.Format(time.RFC3339)
		st := autobanSnapshot(nil)
		res.Status = &st
		recheck429Mu.Lock()
		recheck429Last = res
		recheck429Mu.Unlock()
		return res, nil
	}

	authResp, err := callHostAuthList()
	if err != nil {
		res.Running = false
		res.Message = "host.auth.list: " + err.Error()
		res.Errors = 1
		res.LastRun = now.Format(time.RFC3339)
		recheck429Mu.Lock()
		recheck429Last = res
		recheck429Mu.Unlock()
		return res, err
	}

	// Index xAI auth files by common id keys for O(1) match.
	byKey := map[string]authFile{}
	byEmail := map[string]authFile{}
	for _, f := range authResp.Files {
		if !isXAIAuth(f) {
			continue
		}
		for _, key := range []string{f.ID, f.AuthIndex, f.Name, f.Path, filepathBase(f.Path)} {
			key = strings.TrimSpace(key)
			if key != "" {
				byKey[key] = f
			}
		}
		em := strings.ToLower(strings.TrimSpace(firstNonEmpty(f.Email, f.Account, f.Label)))
		if em != "" {
			byEmail[em] = f
		}
	}

	req := scanRequest{
		Workers:       8,
		TimeoutSec:    20,
		Model:         "grok-4.5",
		Prompt:        "ping",
		ClientVersion: "0.2.93",
		MaxOutputTok:  1,
	}
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        16,
			MaxIdleConnsPerHost: 16,
			Proxy:               http.ProxyFromEnvironment,
		},
		Timeout: time.Duration(req.TimeoutSec) * time.Second,
	}
	ctx := context.Background()

	// Deduplicate by resolved auth file path/index so multi-alias bans probe once.
	type workItem struct {
		banIDs []string
		file   authFile
		email  string
	}
	seenFile := map[string]int{} // file key → work index
	var work []workItem
	for _, t := range targets {
		file, ok := byKey[t.id]
		if !ok {
			em := strings.ToLower(strings.TrimSpace(t.entry.Email))
			if em != "" {
				file, ok = byEmail[em]
			}
		}
		if !ok {
			// Auth gone: drop sticky ban so it cannot block forever.
			runtimeBans.clear(t.id)
			if em := strings.TrimSpace(t.entry.Email); em != "" {
				runtimeBans.clearByEmail(em)
			}
			res.Unbanned++
			res.Details = append(res.Details, recheck429Item{
				AuthID: t.id, Email: t.entry.Email, Action: "unbanned", Detail: "auth file not found",
			})
			continue
		}
		fk := firstNonEmpty(file.AuthIndex, file.ID, file.Path, file.Name)
		if idx, exists := seenFile[fk]; exists {
			work[idx].banIDs = append(work[idx].banIDs, t.id)
			continue
		}
		seenFile[fk] = len(work)
		work = append(work, workItem{banIDs: []string{t.id}, file: file, email: t.entry.Email})
	}

	const workers = 8
	type probeOut struct {
		item workItem
		res  probeResult
	}
	jobs := make(chan workItem)
	outs := make(chan probeOut, len(work))
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for it := range jobs {
				outs <- probeOut{item: it, res: probeOne(ctx, client, it.file, req)}
			}
		}()
	}
	go func() {
		for _, it := range work {
			jobs <- it
		}
		close(jobs)
		wg.Wait()
		close(outs)
	}()

	for po := range outs {
		pr := po.res
		ids := po.item.banIDs
		email := firstNonEmpty(pr.Email, po.item.email)
		// Network / load failure: keep isolation (don't falsely unban).
		if pr.HTTPStatus == 0 || pr.Action == "ERROR" || (pr.Error != "" && pr.HTTPStatus == 0) {
			res.Skipped++
			res.Details = append(res.Details, recheck429Item{
				AuthID: ids[0], Email: email, HTTPStatus: pr.HTTPStatus,
				Action: "skipped", Detail: firstNonEmpty(pr.Error, pr.Summary, "probe failed"),
			})
			continue
		}
		if pr.HTTPStatus == http.StatusTooManyRequests {
			// Still limited: refresh 2h window under all aliases.
			entry, ok := classifyFailure(pr.HTTPStatus, nil, time.Now(), email)
			if ok {
				entry.Source = "recheck429"
				entry.Email = email
				entry.Label = firstNonEmpty(pr.Name, pr.File)
				for _, id := range ids {
					runtimeBans.set(id, entry)
				}
				// Also set canonical ids from probe.
				for _, id := range []string{pr.AuthID, pr.AuthIndex, pr.Name, pr.File} {
					if strings.TrimSpace(id) != "" {
						runtimeBans.set(id, entry)
					}
				}
			}
			res.Still429++
			res.Details = append(res.Details, recheck429Item{
				AuthID: ids[0], Email: email, HTTPStatus: 429,
				Action: "still_429", Detail: "refreshed +2h",
			})
			continue
		}

		// Not 429 → unban all related ids / email.
		for _, id := range ids {
			runtimeBans.clear(id)
		}
		if email != "" {
			runtimeBans.clearByEmail(email)
		}
		for _, id := range []string{pr.AuthID, pr.AuthIndex, pr.Name, pr.File} {
			if strings.TrimSpace(id) != "" {
				runtimeBans.clear(id)
			}
		}

		// If now 401/402/403, re-isolate under correct policy.
		if pr.HTTPStatus == 401 || pr.HTTPStatus == 402 || pr.HTTPStatus == 403 {
			noteScanBan(pr)
			res.Reclassified++
			res.Details = append(res.Details, recheck429Item{
				AuthID: ids[0], Email: email, HTTPStatus: pr.HTTPStatus,
				Action: "reclassified", Detail: fmt.Sprintf("was 429 → now %d", pr.HTTPStatus),
			})
		} else {
			res.Unbanned++
			res.Details = append(res.Details, recheck429Item{
				AuthID: ids[0], Email: email, HTTPStatus: pr.HTTPStatus,
				Action: "unbanned", Detail: "no longer 429",
			})
		}
	}

	res.Running = false
	res.LastRun = time.Now().Format(time.RFC3339)
	res.Message = fmt.Sprintf(
		"%s checked=%d still_429=%d unbanned=%d reclassified=%d skipped=%d",
		mode, res.Checked, res.Still429, res.Unbanned, res.Reclassified, res.Skipped,
	)

	recheck429Mu.Lock()
	// Cap stored details for auto runs.
	stored := res
	if !manual && len(stored.Details) > 40 {
		stored.Details = stored.Details[:40]
	}
	recheck429Last = stored
	recheck429Mu.Unlock()

	// Don't embed full ban list in recheck response (use /bans?page=).
	res.Status = nil
	return res, nil
}

func handleAutobanUnban(body []byte, query url.Values) ([]byte, error) {
	var req struct {
		AuthID  string   `json:"auth_id"`
		AuthIDs []string `json:"auth_ids"`
		Status  int      `json:"status"`
		Email   string   `json:"email"`
		All     bool     `json:"all"`
	}
	if len(body) > 0 {
		_ = json.Unmarshal(body, &req)
	}
	if query != nil {
		if req.AuthID == "" {
			req.AuthID = query.Get("auth_id")
		}
		if req.Status == 0 {
			if s, err := strconv.Atoi(query.Get("status")); err == nil {
				req.Status = s
			}
		}
		if !req.All && query.Get("all") == "1" {
			req.All = true
		}
		if len(req.AuthIDs) == 0 {
			if raw := query.Get("auth_ids"); raw != "" {
				req.AuthIDs = strings.Split(raw, ",")
			}
		}
		if req.Email == "" {
			req.Email = query.Get("email")
		}
	}

	removed := 0
	switch {
	case req.All:
		removed = runtimeBans.clearAll()
	case req.Status > 0:
		removed = runtimeBans.clearStatus(req.Status)
	case req.Email != "":
		removed = runtimeBans.clearByEmail(req.Email)
	case len(req.AuthIDs) > 0:
		removed = runtimeBans.clearMany(req.AuthIDs)
	case strings.TrimSpace(req.AuthID) != "":
		if runtimeBans.clear(req.AuthID) {
			removed = 1
		}
	default:
		return jsonErrorEnvelope(http.StatusBadRequest, "bad_request", "provide auth_id, auth_ids, status, email, or all=true")
	}
	return jsonManagementEnvelope(http.StatusOK, map[string]any{
		"ok": true, "removed": removed, "status": autobanSnapshot(nil),
	})
}

func handleAutobanImport(body []byte) ([]byte, error) {
	var snapshot autobanStatus
	if err := json.Unmarshal(body, &snapshot); err != nil {
		return jsonErrorEnvelope(http.StatusBadRequest, "bad_request", err.Error())
	}
	now := time.Now()
	imported := 0
	for _, item := range snapshot.Bans {
		resetAt, errReset := time.Parse(time.RFC3339, item.ResetAt)
		if errReset != nil || !resetAt.After(now) || strings.TrimSpace(item.AuthID) == "" {
			continue
		}
		bannedAt, errBanned := time.Parse(time.RFC3339, item.BannedAt)
		if errBanned != nil {
			bannedAt = now
		}
		runtimeBans.set(item.AuthID, banEntry{
			StatusCode: item.StatusCode,
			Reason:     firstNonEmpty(item.Reason, "import"),
			Source:     "import",
			Email:      item.Email,
			Label:      item.Label,
			BannedAt:   bannedAt,
			ResetAt:    resetAt,
			FailCount:  item.FailCount,
		})
		imported++
	}
	return jsonManagementEnvelope(http.StatusOK, map[string]any{
		"ok": true, "imported": imported, "status": autobanSnapshot(nil),
	})
}
