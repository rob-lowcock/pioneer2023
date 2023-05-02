package handlers

import (
	"log"
	"net/http"

	"github.com/google/jsonapi"
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

	requestModel := new(models.Auth)

	err := jsonapi.UnmarshalPayload(r.Body, requestModel)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.Auth.ValidateCredentials(requestModel.Email, requestModel.Password)
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

	tokenModel := new(models.JWT)
	tokenModel.ID = token

	err = jsonapi.MarshalPayload(w, tokenModel)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
