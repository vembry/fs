package handler

import (
	"api/pkg/db"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func (h *handler) Test(w http.ResponseWriter, r *http.Request) {
	constructResponse(w, http.StatusOK, map[string]interface{}{
		"message": "ok",
	})
}

func (h *handler) InsertTestPath(w http.ResponseWriter, r *http.Request) {

	raw, _ := json.Marshal(map[string]interface{}{
		"field1": "some-value",
	})

	err := h.q.InsertPath(r.Context(), db.InsertPathParams{
		Path:        fmt.Sprintf("dummy-path-%d", time.Now().UnixMilli()),
		Information: raw,
	})
	if err != nil {
		slog.Error("error on inserting dummy path", slog.String("error", err.Error()))

		constructResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "failed on creating dummy entry",
			"error":   err.Error(),
		})
		return
	}

	constructResponse(w, http.StatusOK, struct{}{})
}

func constructResponse(w http.ResponseWriter, httpStatusCode int, body interface{}) {
	raw, _ := json.Marshal(body)
	w.WriteHeader(httpStatusCode)
	w.Write(raw)
}
