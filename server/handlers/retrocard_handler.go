package handlers

import (
	"log"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/rob-lowcock/pioneer2023/db"
	"github.com/rob-lowcock/pioneer2023/models"
)

type GetRetrocardHandler struct {
	RetrocardDb db.Retrocard
}

func (h GetRetrocardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cards, err := h.RetrocardDb.GetActiveCards()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Convert the cards to a slice of pointers for JSONAPI
	cardsPtr := []*models.Retrocard{}
	for i := range cards {
		cardsPtr = append(cardsPtr, &cards[i])
	}

	err = jsonapi.MarshalPayload(w, cardsPtr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
