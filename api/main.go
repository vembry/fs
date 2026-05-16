package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	// handler
	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := json.Marshal(map[string]interface{}{
			"message":  "ok",
			"message1": "ok1",
		})
		w.WriteHeader(http.StatusOK)
		w.Write(raw)
	})

	mux.HandleFunc("GET /f", func(w http.ResponseWriter, r *http.Request) {
		entries, err := os.ReadDir("/mounts/sample directory")
		if err != nil {
			raw, _ := json.Marshal(map[string]interface{}{
				"message": "failed to browse directory",
				"error":   err.Error(),
			})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(raw)
			return
		}

		entriesOut := []string{}
		for _, entry := range entries {
			fmt.Println("Name:", entry.Name())
			fmt.Println("Is Directory:", entry.IsDir())
			fmt.Println("-----------------")
			entriesOut = append(entriesOut, entry.Name())
		}

		raw, _ := json.Marshal(map[string]interface{}{
			// "path":    pathQuery,
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
