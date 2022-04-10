package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"petstore/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// TestHandler_Post is used to test Post method
func TestHandler_Post(t *testing.T) {
	id := uuid.New()
	pet := models.Pet{ID: id, Name: "bruno", Age: 2, Species: "dog", Status: "Available"}
	testcases := []struct {
		desc   string
		input  models.Pet
		output models.Pet
		err    error
		status int
	}{
		{"success case", pet, pet, nil, http.StatusCreated},
	}
	ctrl := gomock.NewController(t)
	MockService := NewMockService(ctrl)
	s := New(MockService)
	ctx := context.Background()

	for _, tc := range testcases {
		body, err := json.Marshal(tc.input)
		if err != nil {
			return
		}

		req := httptest.NewRequest(http.MethodPost, "/pet", bytes.NewBuffer(body))

		w := httptest.NewRecorder()

		MockService.EXPECT().ServicePost(ctx, pet).Return(tc.output, tc.err)
		s.Post(w, req)
		assert.Equal(t, tc.status, w.Code)
	}
}

func TestHandler_PostSerErr(t *testing.T) {
	id := uuid.New()
	pet := models.Pet{ID: id, Name: "bruno", Age: 2, Species: "dog", Status: "Available"}
	testcases := []struct {
		desc   string
		input  models.Pet
		output models.Pet
		err    error
		status int
	}{
		{"success case", pet, models.Pet{}, errors.New("error from service layer"), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	MockService := NewMockService(ctrl)
	s := New(MockService)
	ctx := context.Background()

	for _, tc := range testcases {
		body, err := json.Marshal(tc.input)
		if err != nil {
			return
		}

		req := httptest.NewRequest(http.MethodPost, "/pet", bytes.NewBuffer(body))

		w := httptest.NewRecorder()

		MockService.EXPECT().ServicePost(ctx, pet).Return(tc.output, tc.err)
		s.Post(w, req)
		assert.Equal(t, tc.status, w.Code)
	}
}

func Test_PostUnMarErr(t *testing.T) {
	pet := []byte(`{"ID": id, "Name": "bruno", "Age": 2, "Species": "dog", "Status": "Available"}`)

	testcases := []struct {
		desc   string
		input  []byte
		err    error
		status int
	}{
		{"success", pet, nil, http.StatusBadRequest},
	}

	ctrl := gomock.NewController(t)
	MockService := NewMockService(ctrl)
	s := New(MockService)

	for _, tc := range testcases {
		body, _ := json.Marshal(tc.input)
		r := httptest.NewRequest(http.MethodPost, "/pets", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		s.Post(w, r)
		assert.Equal(t, tc.status, w.Code)
	}
}

// TestHandler_GetByID is used to test the GetByID method in Handlers layer
func TestHandler_GetByID(t *testing.T) {
	id := uuid.New()
	pet := models.Pet{ID: id, Name: "bruno", Age: 2, Species: "dog", Status: "Available"}
	testcases := []struct {
		desc   string
		input  uuid.UUID
		output models.Pet
		err    error
		status int
	}{
		{"success", id, pet, nil, http.StatusOK},
	}
	ctrl := gomock.NewController(t)
	MockService := NewMockService(ctrl)
	s := New(MockService)

	for _, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/pet/{id}", nil)
		r := mux.SetURLVars(req, map[string]string{"id": tc.input.String()})

		MockService.EXPECT().ServiceGetByID(gomock.Any(), tc.input).Return(tc.output, tc.err)
		s.GetByID(w, r)
		assert.Equal(t, tc.status, w.Code)
	}
}

func TestHandler_GetByIDSerErr(t *testing.T) {
	id := uuid.New()
	testcases := []struct {
		desc   string
		input  uuid.UUID
		output models.Pet
		err    error
		status int
	}{
		{"success", id, models.Pet{}, errors.New("error from service layer"), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	MockService := NewMockService(ctrl)
	s := New(MockService)

	for _, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/pet/{id}", nil)
		r := mux.SetURLVars(req, map[string]string{"id": tc.input.String()})

		MockService.EXPECT().ServiceGetByID(gomock.Any(), tc.input).Return(tc.output, tc.err)
		s.GetByID(w, r)
		assert.Equal(t, tc.status, w.Code)
	}
}
func TestHandler_GetByIDInvalidID(t *testing.T) {

	testcases := []struct {
		desc   string
		input  string
		output models.Pet
		err    error
		status int
	}{
		{"success", "1", models.Pet{}, nil, http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	MockService := NewMockService(ctrl)
	s := New(MockService)

	for _, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/pet/{id}", nil)
		r := mux.SetURLVars(req, map[string]string{"id": tc.input})

		s.GetByID(w, r)
		assert.Equal(t, tc.status, w.Code)
	}
}
