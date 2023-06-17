package retrocard

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/jsonapi"
	"github.com/rob-lowcock/pioneer2023/db"
	"github.com/rob-lowcock/pioneer2023/models"
)

type UpdateRetrocardHandler struct {
	RetrocardDb db.Retrocard
}

func (h UpdateRetrocardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	requestModel := new(models.Retrocard)
	err := jsonapi.UnmarshalPayload(r.Body, requestModel)

	if chi.URLParam(r, "id") != requestModel.ID {
		log.Println("id mismatch in update retrocard request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.RetrocardDb.UpdateRetrocard(*requestModel)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = jsonapi.MarshalPayload(w, requestModel)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
