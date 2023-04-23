package handlers

import "net/http"

type LoginHandler struct {
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the login page! Have I got it?"))
}
