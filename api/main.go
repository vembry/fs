package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type PathEntry struct {
	Path        string `json:"path"`
	IsDirectory bool   `json:"is_directory"`
}

func main() {
	mux := http.NewServeMux()

	// handler
	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := json.Marshal(map[string]interface{}{
			"message": "ok",
		})
		w.WriteHeader(http.StatusOK)
		w.Write(raw)
	})

	mux.HandleFunc("GET /f", func(w http.ResponseWriter, r *http.Request) {
		pathQuery := r.URL.Query().Get("d")

		fullPath := "/mounts"
		paths := []string{}
		if len(strings.TrimSpace(pathQuery)) > 0 {
			fullPath = fmt.Sprintf("%s/%s", fullPath, pathQuery)
			paths = strings.Split(pathQuery, "/")
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
		entriesOut := []PathEntry{}
		for _, entry := range entries {
			entriesOut = append(entriesOut, PathEntry{
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
	})

	// construct server
	httpserver := http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	log.SetOutput(os.Stdout)
	log.Printf("starting server...")
	// start server
	if err := httpserver.ListenAndServe(); err != nil {
		log.Fatalf("error on starting server. err=%v", err)
	}
	log.Printf("shutting down server...")
}
