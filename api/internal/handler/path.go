package handler

import (
	"log/slog"
	"net/http"
)

func (h *handler) ListPaths(w http.ResponseWriter, r *http.Request) {
	paths, err := h.q.ListPaths(r.Context())
	if err != nil {
		slog.Error("error on listing out paths", slog.String("error", err.Error()))
		constructResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	constructResponse(w, http.StatusOK, map[string]interface{}{
		"paths": paths,
	})
}
