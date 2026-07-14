package main

import (
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

// Public release stubs: no SSO convert / vault. Full features live in grok-manager-cpa.

var (
	vaultMu   sync.Mutex
	vaultPath string
)

func pluginDataCandidates(fileName string) []string {
	var out []string
	if wd, err := os.Getwd(); err == nil {
		out = append(out,
			filepath.Join(wd, "plugins", "grok-manager", fileName),
			filepath.Join(wd, "plugins-data", "grok-manager", fileName),
		)
	}
	out = append(out,
		filepath.Join(`C:\CLIProxyAPI-local\plugins\grok-manager`, fileName),
		filepath.Join(`/root/.cli-proxy-api/plugins/grok-manager`, fileName),
	)
	return out
}

func resolvePluginDataPath(fileName string, cached *string) string {
	if cached != nil && *cached != "" {
		return *cached
	}
	for _, p := range pluginDataCandidates(fileName) {
		if st, err := os.Stat(filepath.Dir(p)); err == nil && st.IsDir() {
			if cached != nil {
				*cached = p
			}
			return p
		}
	}
	p := pluginDataCandidates(fileName)[0]
	_ = os.MkdirAll(filepath.Dir(p), 0o755)
	if cached != nil {
		*cached = p
	}
	return p
}

func vaultEmailSet() map[string]bool {
	return map[string]bool{}
}

func vaultMeta() (path string, count int, savedAt string) {
	return "", 0, ""
}

func vaultLookupByEmail(email string) (ssoVaultEntry, bool) {
	return ssoVaultEntry{}, false
}

func loadSSOHistoryOnStart() {}

func syncVaultHTTPFromScan(results []probeResult) {}

func autoRefresh401FromVault(workers int) {}

func noteSSOSuccess(email, file string) {}

type ssoVaultEntry struct {
	Email string
	SSO   string
}

func vaultPublicSummary(q url.Values) map[string]any {
	return map[string]any{
		"count": 0, "entries": []any{}, "page": 1, "pages": 1, "match": 0,
		"disabled": true,
	}
}
