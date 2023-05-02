package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthHandler struct {
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	out, _ := json.Marshal(map[string]interface{}{
		"status": "ok",
	})
	w.Write(out)
}
