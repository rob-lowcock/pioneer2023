package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/google/jsonapi"
	"github.com/rob-lowcock/pioneer2023/auth"
	"github.com/rob-lowcock/pioneer2023/models"
)

type LoginHandler struct {
	Auth auth.Auth
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", os.Getenv("CLIENT_URL"))
	w.Header().Add("Vary", "Origin")

	if r.Method == "OPTIONS" {
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Add("Content-Type", "application/vnd.api+json")

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
