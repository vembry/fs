package main

import (
	"encoding/json"
	"log"
	"net/http"
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

	// construct server
	httpserver := http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	// start server
	if err := httpserver.ListenAndServe(); err != nil {
		log.Fatalf("error on starting server. err=%v", err)
	}
}
