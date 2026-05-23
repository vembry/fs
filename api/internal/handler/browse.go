package handler

import (
	"api/internal/model"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func isUnsafePathing(path string) bool {
	return strings.Contains(path, "..")
}

func (h *handler) Browse(w http.ResponseWriter, r *http.Request) {
	pathQuery := r.URL.Query().Get("d")
	slog.Info("logging path-query", slog.String("pathQuery", pathQuery))

	// checks for possible directory traversal attack
	if isUnsafePathing(pathQuery) {
		slog.InfoContext(r.Context(), "caught attempt to use unsafe path", slog.String("pathQuery", pathQuery))

		raw, _ := json.Marshal(map[string]interface{}{
			"message": "'dot-dot-slash' in 'd' parameter is not allowed",
			"path":    pathQuery,
		})
		w.WriteHeader(http.StatusForbidden)
		w.Write(raw)
		return
	}

	// clean build the directory path
	fullPath := "/mounts"
	if len(strings.TrimSpace(pathQuery)) > 0 {
		fullPath = filepath.Join("/mounts", pathQuery)
		slog.Info("logging full-path", slog.String("fullPath", fullPath))
	}

	// construct the path references for the response
	temp := strings.TrimLeft(fullPath, "/mounts")
	paths := []string{}
	if len(temp) > 0 {
		paths = strings.Split(temp, "/")
	}

	// browse directory
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		raw, _ := json.Marshal(map[string]interface{}{
			"message": "failed to browse directory",
			"error":   err.Error(),
		})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(raw)
		return
	}

	// construct response
	entriesOut := []model.PathEntry{}
	for _, entry := range entries {
		slog.Info("logging file entry", slog.Any("entry", entry))
		entriesOut = append(entriesOut, model.PathEntry{
			Path:        entry.Name(),
			IsDirectory: entry.IsDir(),
		})
	}

	raw, _ := json.Marshal(map[string]interface{}{
		"paths":   paths,
		"entries": entriesOut,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}
