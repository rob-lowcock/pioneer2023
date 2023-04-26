package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rob-lowcock/pioneer2023/auth"
	"github.com/rob-lowcock/pioneer2023/models"
)

type LoginHandler struct {
	Auth auth.Auth
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user, err := h.Auth.ValidateCredentials(r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := h.Auth.GenerateToken(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(models.JWT{Token: token})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(output)
}
