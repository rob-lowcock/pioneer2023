package retrocard

import (
	"net/http"

	"github.com/rob-lowcock/pioneer2023/db"
)

type RetrocardHandler struct {
	RetrocardDb db.Retrocard
}

func (h RetrocardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handler := GetRetrocardHandler{h.RetrocardDb}
		handler.ServeHTTP(w, r)
		return
	}

	if r.Method == "POST" {
		handler := CreateRetrocardHandler{h.RetrocardDb}
		handler.ServeHTTP(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
