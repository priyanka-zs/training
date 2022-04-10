package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"petstore/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	service Service
}

func New(s Service) Handler {
	return Handler{service: s}
}

// Post method is used to take http request and respond accordingly
func (h Handler) Post(w http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	w.Header().Set("Context-Type", "application/json")

	body, err := io.ReadAll(request.Body)
	if err != nil {
		w.Write([]byte("io.read all err"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var pet models.Pet

	err = json.Unmarshal(body, &pet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.service.ServicePost(ctx, pet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(result)
	if err != nil {
		return
	}
}

// GetByID takes id as input and displays a row of pet
func (h Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Context-Type", "application/json")

	paraID := mux.Vars(r)
	Pid := paraID["id"]

	ID, err := uuid.Parse(Pid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pet, err := h.service.ServiceGetByID(ctx, ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(pet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(res)
	if err != nil {
		return
	}
}
