package main

import (
	"api/internal/handler"
	"api/pkg/db"
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type httpserver struct {
	s *http.Server
}

func newHttpServer(q db.Querier) *httpserver {
	mux := http.NewServeMux()
	h := handler.NewHandler(q)

	// handler
	mux.HandleFunc("GET /test", h.Test)
	mux.HandleFunc("POST /test/path", h.InsertTestPath)
	mux.HandleFunc("GET /f", h.Browse)
	mux.HandleFunc("GET /path", h.ListPaths)

	return &httpserver{
		s: &http.Server{
			Addr:    ":80",
			Handler: mux,
		},
	}
}

func (h *httpserver) Start() {
	go func() {
		// start server
		if err := h.s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error on starting server. err=%v", err)
		}
	}()
}

func (h *httpserver) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := h.s.Shutdown(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "error on shutting down http-server", slog.String("error", err.Error()))
	}
	return nil
}
