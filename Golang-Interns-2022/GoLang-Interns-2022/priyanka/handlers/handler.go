package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/zopsmart/GoLang-Interns-2022/model"
)

type Handler struct {
	service Car
}

func New(car Car) Handler {
	return Handler{service: car}
}

// GetByBrand method is used in handler layer to fetch rows.
func (h Handler) GetByBrand(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	query1 := request.URL.Query().Get("brand")
	query2 := request.URL.Query().Get("isEngine")

	isEngine, err := strconv.ParseBool(query2)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := h.service.Get(ctx, query1, isEngine)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}

		return
	}

	body, err := json.Marshal(res)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	writer.Header().Set("Content-Type", "application-json")

	_, err = writer.Write(body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		return
	}
}

// GetByID method is used in handler layer to fetch row.
func (h Handler) GetByID(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	writer.Header().Set("Content-Type", "application/json")

	param := mux.Vars(request)

	id := param["id"]

	uuid1, err := uuid.Parse(id)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}

		return
	}

	resp, err := h.service.GetByID(ctx, uuid1)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	body, err := json.Marshal(resp)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	_, err = writer.Write(body)
	if err != nil {
		return
	}
}

// Update method is used in handler layer to update rows.
func (h Handler) Update(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var car model.Car

	body, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &car)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	para := mux.Vars(request)

	var paraID = para["id"]

	id, err := uuid.Parse(paraID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	car.ID = id

	res, err := h.service.Update(ctx, &car)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}

		return
	}

	body, err = json.Marshal(res)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	writer.Header().Set("Content-type", "application/json")

	_, err = writer.Write(body)
	if err != nil {
		return
	}
}

// Create method is used in handler layer for inserting rows.
func (h Handler) Create(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	writer.Header().Set("Content-Type", "application/json")

	var car model.Car
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &car)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.service.Create(ctx, &car)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}

		return
	}

	result, err := json.Marshal(resp)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	writer.WriteHeader(http.StatusCreated)

	_, err = writer.Write(result)
	if err != nil {
		return
	}
}

// Delete method is used in handler layer for deleting rows.
func (h Handler) Delete(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	writer.Header().Set("Content-Type", "application/json")

	param := mux.Vars(request)
	id := param["id"]

	uuid1, err := uuid.Parse(id) // string to uuid conversion.

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.Delete(ctx, uuid1)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	writer.WriteHeader(http.StatusNoContent)
}
