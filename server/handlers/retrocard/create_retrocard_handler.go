package retrocard

import (
	"log"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/rob-lowcock/pioneer2023/db"
	"github.com/rob-lowcock/pioneer2023/models"
)

type CreateRetrocardHandler struct {
	RetrocardDb db.Retrocard
}

func (h CreateRetrocardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	requestModel := new(models.Retrocard)

	err := jsonapi.UnmarshalPayload(r.Body, requestModel)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.RetrocardDb.CreateRetrocard(*requestModel)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	requestModel.ID = id

	w.WriteHeader(http.StatusCreated)
	err = jsonapi.MarshalPayload(w, requestModel)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
