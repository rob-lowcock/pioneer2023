package handlers

import (
	"net/http"

	"github.com/go-oauth2/oauth2/v4/server"
)

type LoginHandler struct {
	AuthServer *server.Server
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.AuthServer.HandleTokenRequest(w, r)
}
