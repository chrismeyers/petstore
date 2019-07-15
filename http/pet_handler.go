package http

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/chrismeyers/petstore"
	"github.com/julienschmidt/httprouter"
)

type PetHandler struct {
	*httprouter.Router
	Service petstore.PetService
	Logger  *log.Logger
}

func NewPetHandler(p petstore.PetService) *PetHandler {
	logger := log.New(os.Stderr, "", log.LstdFlags)

	h := &PetHandler{
		Router:  httprouter.New(),
		Service: p,
		Logger:  logger,
	}

	h.GET("/pets", h.handleGetPets())
	h.GET("/pets/:id", h.handleGetPet())
	h.POST("/pets", h.handleAddPet())
	h.DELETE("/pets/:id", h.handleDeletePet())

	return h
}

func (h *PetHandler) handleGetPets() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		pet, err := h.Service.GetAll()

		if err != nil {
			Error(w, err, http.StatusInternalServerError, h.Logger)
			return
		}

		EncodeJSON(w, pet, h.Logger)
	}
}

func (h *PetHandler) handleGetPet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		iid, err := strconv.Atoi(id)

		if err != nil {
			Error(w, err, http.StatusBadRequest, h.Logger)
			return
		}

		pet, err := h.Service.Get(iid)

		if err != nil {
			Error(w, err, http.StatusInternalServerError, h.Logger)
		}

		EncodeJSON(w, pet, h.Logger)
	}
}

func (h *PetHandler) handleAddPet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var body petstore.Pet
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Error(w, err, http.StatusBadRequest, h.Logger)
			return
		}

		id, err := h.Service.Create(body)

		if err != nil {
			Error(w, err, http.StatusInternalServerError, h.Logger)
		}

		EncodeJSON(w, map[string]int{"id": id}, h.Logger)
	}
}

func (h *PetHandler) handleDeletePet() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		iid, err := strconv.Atoi(id)

		if err != nil {
			Error(w, err, http.StatusBadRequest, h.Logger)
			return
		}

		err = h.Service.Delete(iid)

		if err != nil {
			Error(w, err, http.StatusInternalServerError, h.Logger)
		}

		EncodeJSON(w, nil, h.Logger)
	}
}
