package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func handlePathsInfo() ([]byte, error) {
	job.mu.Lock()
	hist := job.historyPath
	job.mu.Unlock()
	if hist == "" {
		hist = resolvePluginDataPath(historyFileName, nil)
	}
	sch := scheduleSnapshot()
	wd, _ := os.Getwd()
	return jsonManagementEnvelope(http.StatusOK, map[string]any{
		"plugin":       pluginName,
		"version":      pluginVersion,
		"scan_history": hist,
		"schedule_path": sch.ConfigPath,
		"bans_path":    bansFilePath(),
		"bans_count":   runtimeBans.count(),
		"wd":           wd,
		"note":         "public build: scan / autoban / schedule only",
	})
}

func handleBackup() ([]byte, error) {
	files := []struct {
		name string
		path string
	}{
		{"last-scan.json", resolvePluginDataPath(historyFileName, nil)},
		{"schedule.json", resolvePluginDataPath(scheduleFileName, nil)},
		{"bans.json", bansFilePath()},
	}
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	added := 0
	for _, f := range files {
		raw, err := os.ReadFile(f.path)
		if err != nil || len(raw) == 0 {
			continue
		}
		w, err := zw.Create(f.name)
		if err != nil {
			continue
		}
		if _, err := w.Write(raw); err != nil {
			continue
		}
		added++
	}
	_ = zw.Close()
	if added == 0 {
		return jsonErrorEnvelope(http.StatusBadRequest, "empty", "nothing to backup")
	}
	name := fmt.Sprintf("backup-%s.zip", time.Now().UTC().Format("20060102-150405"))
	outPath := resolvePluginDataPath(name, nil)
	if err := os.WriteFile(outPath, buf.Bytes(), 0o644); err != nil {
		return jsonErrorEnvelope(http.StatusInternalServerError, "write_failed", err.Error())
	}
	return jsonManagementEnvelope(http.StatusOK, map[string]any{
		"ok": true, "path": outPath, "filename": name, "bytes": buf.Len(), "files": added,
	})
}

// silence unused import if filepath only used above
var _ = filepath.Separator
