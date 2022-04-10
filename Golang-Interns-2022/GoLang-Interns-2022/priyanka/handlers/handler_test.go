package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/zopsmart/GoLang-Interns-2022/model"
)

// TestHandler_Delete is used to test Delete method in handler layer.
func TestHandler_Delete(t *testing.T) {
	ID := uuid.New()
	testcases := []struct {
		desc     string
		id       uuid.UUID
		expected int
		err      error
	}{
		{"success case", ID, http.StatusNoContent, nil},
		{"id not given", uuid.Nil, http.StatusBadRequest, errors.New("invalid id")},
	}

	ctrl := gomock.NewController(t)

	m := NewMockCar(ctrl)
	h := New(m)

	for i, tc := range testcases {
		req := httptest.NewRequest(http.MethodDelete, "https://car", nil)
		r := mux.SetURLVars(req, map[string]string{"id": tc.id.String()})
		w := httptest.NewRecorder()

		gomock.InOrder(m.EXPECT().Delete(gomock.Any(), tc.id).Return(tc.err))
		h.Delete(w, r)

		if w.Code != tc.expected {
			t.Errorf("[Test %d]Failed. Got %v Expected %v/n", i+1, w.Code, tc.expected)
		}
	}
}

// TestHandler_DeleteError is used to test delete error case
func TestHandler_DeleteError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockCar(ctrl)

	testcases := []struct {
		desc   string
		id     string
		status int
	}{
		{"id invalid", "123", http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest(http.MethodDelete, "https://car/id", nil)
		w := httptest.NewRecorder()

		h := New(m)
		h.Delete(w, req)

		assert.Equal(t, tc.status, w.Code)
	}
}

// TestHandler_GetById is used to test GetById method in handler layer.
func TestHandler_GetByID(t *testing.T) {
	ID := uuid.New()
	ID1 := uuid.New()
	var car1 = model.Car{
		ID: ID, Name: "bmw3", Year: 2005, Brand: "bmw",
		Fuel: "petrol", Engine: model.Engine{Displacement: 150, Ranges: 0, NoOfCylinders: 3,
			EngineID: ID1}}

	testcases := []struct {
		desc           string
		id             uuid.UUID
		resBody        *model.Car
		err            error
		expectedStatus int
	}{
		{"success case", ID, &car1, nil, http.StatusOK},
		{"id not given", uuid.Nil, &model.Car{}, errors.New("error"), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	m := NewMockCar(ctrl)
	h := New(m)

	for i, tc := range testcases {
		req := httptest.NewRequest(http.MethodGet, "https://car/id", nil)
		r := mux.SetURLVars(req, map[string]string{"id": tc.id.String()})
		w := httptest.NewRecorder()

		gomock.InOrder(m.EXPECT().GetByID(gomock.Any(), tc.id).Return(tc.resBody, tc.err))
		h.GetByID(w, r)

		if w.Code != tc.expectedStatus {
			t.Errorf("[Test %d]Failed Got %v expected %v", i+1, w.Code, tc.expectedStatus)
		}
	}
}

// TestHandler_GetByIDErr is used to test GetByID
func TestHandler_GetByIDErr(t *testing.T) {
	testcases := []struct {
		desc           string
		id             string
		resBody        *model.Car
		err            error
		expectedStatus int
	}{
		{"id not given", "123", &model.Car{}, errors.New("invalid uuid"), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	m := NewMockCar(ctrl)
	h := New(m)

	for i, tc := range testcases {
		req := httptest.NewRequest(http.MethodGet, "https://car/id", nil)
		r := mux.SetURLVars(req, map[string]string{"id": tc.id})
		w := httptest.NewRecorder()

		h.GetByID(w, r)

		if w.Code != tc.expectedStatus {
			t.Errorf("[Test %d]Failed Got %v expected %v", i+1, w.Code, tc.expectedStatus)
		}
	}
}

// TestHandler_GetByBrand is used to test GetByBrand method in handler layer.
func TestHandler_GetByBrand(t *testing.T) {
	ID := uuid.New()
	ID1 := uuid.New()
	var car1 = model.Car{ID: ID, Name: "b3", Year: 2003,
		Brand: "bmw", Fuel: "petrol", Engine: model.Engine{Displacement: 200, Ranges: 0, NoOfCylinders: 3,
			EngineID: ID1}}

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockCar(ctrl)

	testcases := []struct {
		desc           string
		brand          string
		engine         bool
		resBody        []*model.Car
		err            error
		expectedStatus int
	}{
		{"success case", "bmw", true, []*model.Car{&car1}, nil, http.StatusOK},
		{"invalid brand", "bm", true, []*model.Car{},
			errors.New("invalid brand"), http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest(http.MethodGet, "https://car?brand="+tc.brand+"&isEngine="+strconv.FormatBool(tc.engine), nil)
		w := httptest.NewRecorder()

		m.EXPECT().Get(gomock.Any(), tc.brand, tc.engine).Return(tc.resBody, tc.err)
		h := New(m)
		h.GetByBrand(w, req)

		assert.Equal(t, w.Code, tc.expectedStatus)
	}
}

// TestHandler_GetByBrandBoolError is used to test GetByBrand
func TestHandler_GetByBrandBoolError(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockCar(ctrl)

	testcases := []struct {
		desc           string
		brand          string
		engine         string
		resBody        []*model.Car
		err            error
		expectedStatus int
	}{
		{"success case", "bmw", "tru", []*model.Car{}, errors.New("error in bool conversion"), http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest(http.MethodGet, "https://car?brand="+tc.brand+"&isEngine="+tc.engine, nil)
		w := httptest.NewRecorder()
		h := New(m)
		h.GetByBrand(w, req)

		assert.Equal(t, w.Code, tc.expectedStatus)
	}
}

// TestHandler_Update is used to test Update in handler
func TestHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockCar(ctrl)
	ID := uuid.New()
	ID1 := uuid.New()
	var car1 = model.Car{
		ID: ID, Name: "bmw3", Year: 2005, Brand: "bmw",
		Fuel: "petrol", Engine: model.Engine{Displacement: 150, Ranges: 0, NoOfCylinders: 3,
			EngineID: ID1}}

	var car2 = model.Car{ID: ID, Name: "bmw3",
		Year: 2005, Brand: "bmw", Fuel: "petrol", Engine: model.Engine{Displacement: 150, Ranges: 0, NoOfCylinders: 3,
			EngineID: ID1}}

	testcases := []struct {
		desc               string
		id                 uuid.UUID
		reqBody            *model.Car
		resBody            *model.Car
		err                error
		expectedStatusCode int
	}{
		{"success case", car1.ID, &car1, &car1, nil, http.StatusOK},
		{"invalid id", car2.ID, &car2, &model.Car{}, errors.New("error from service layer"), http.StatusBadRequest},
	}

	for _, tc := range testcases {
		body, err := json.Marshal(tc.reqBody)
		if err != nil {
			fmt.Println(err)
		}

		req := httptest.NewRequest(http.MethodPut, "https://car/id", bytes.NewBuffer(body))
		r := mux.SetURLVars(req, map[string]string{"id": tc.id.String()})
		w := httptest.NewRecorder()

		m.EXPECT().Update(gomock.Any(), tc.reqBody).Return(tc.resBody, tc.err)
		h := New(m)
		h.Update(w, r)

		assert.Equal(t, w.Code, tc.expectedStatusCode)
	}
}

func TestMockService_CreateError(t *testing.T) {
	testcases := []struct {
		desc   string
		input  []byte
		status int
	}{
		{"missing parameters", []byte(`[{ "Name": "tesla 1t2e",
    "Year": 2016,"Brand": "tesla","Fuel": "diesel","Engine": {"Displacement": 110,"NoOfCylinders": 2}}]`),
			http.StatusBadRequest},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockCar(ctrl)

	for _, tc := range testcases {
		req := httptest.NewRequest(http.MethodPost, "https://car", bytes.NewBuffer(tc.input))
		w := httptest.NewRecorder()

		h := New(m)
		h.Create(w, req)

		assert.Equal(t, tc.status, w.Code)
	}
}

// TestHandler_Create is used to test Create in handler
func TestHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockCar(ctrl)
	ID := uuid.New()
	ID1 := uuid.New()
	var car1 = model.Car{
		ID: ID, Name: "bmw3", Year: 2005, Brand: "bmw",
		Fuel: "petrol", Engine: model.Engine{Displacement: 150, Ranges: 0, NoOfCylinders: 3,
			EngineID: ID1}}

	var car2 = model.Car{ID: ID,
		Name: "bmw3", Year: 2005, Brand: "bmw", Fuel: "petrol",
		Engine: model.Engine{Displacement: 150, Ranges: 0, NoOfCylinders: 3, EngineID: ID1}}

	testcases := []struct {
		desc               string
		reqBody            *model.Car
		resBody            *model.Car
		err                error
		expectedStatusCode int
	}{
		{"success case", &car1, &car1, nil, http.StatusCreated},
		{"invalid id", &car2, &model.Car{}, errors.New("error from service layer"), http.StatusBadRequest},
	}

	for i, tc := range testcases {
		body, err := json.Marshal(tc.reqBody)
		if err != nil {
			return
		}

		req := httptest.NewRequest(http.MethodPost, "https://car", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		m.EXPECT().Create(gomock.Any(), tc.reqBody).Return(tc.resBody, tc.err)
		h := New(m)
		h.Create(w, req)

		if w.Code != tc.expectedStatusCode {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Code, tc.expectedStatusCode)
		}
	}
}
