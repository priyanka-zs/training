package car

import (
	"bytes"
	"manasa/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//8f443772-132b-4ae5-9f8f-9960649b3fb4
var car1 = models.Car{
	Id:    uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
	Name:  "b3",
	Year:  2003,
	Brand: "BMW",
	Fuel:  "petrol",
	Engine: models.Engine{
		Displacement:  200,
		Range:         0,
		NoOfCylinders: 3,
		Id:            uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
	},
}

var car2 = models.Car{
	Id:    uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
	Name:  "t1",
	Year:  2019,
	Brand: "Tesla",
	Fuel:  "electric",
	Engine: models.Engine{
		Displacement:  0,
		Range:         300,
		NoOfCylinders: 0,
		Id:            uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
	},
}

func TestHandler_Delete(t *testing.T) {
	testcases := []struct {
		id       uuid.UUID
		expected int
	}{
		{uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"), http.StatusOK},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockCar(ctrl)

	gomock.InOrder(m.EXPECT().Delete(uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4")).Return(car1, nil))

	h := New(m)

	for i, tc := range testcases {
		req := httptest.NewRequest(http.MethodGet, "https://car", nil)
		r := mux.SetURLVars(req, map[string]string{"id": tc.id.String()})
		w := httptest.NewRecorder()

		h.Delete(w, r)
		if w.Code != tc.expected {
			t.Errorf("[Test %d]Failed. Got %v Expected %v/n", i+1, w.Code, tc.expected)
		}
	}
}

func TestHandler_GetById(t *testing.T) {
	testcases := []struct {
		id             uuid.UUID
		resBody        []byte
		expectedStatus int
	}{
		{uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"), []byte(`{"id":"8f443772-132b-4ae5-9f8f-9960649b3fb4","name":"b3","year":"2003","brand":"BMW","fuel_type":"petrol","engine":{"displacement":200,"range":0,"cylinders":3,"id":"8f443772-132b-4ae5-9f8f-9960649b3fb4"}}`), http.StatusOK},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockCar(ctrl)

	gomock.InOrder(m.EXPECT().GetById(car1.CarId).Return(car1, nil))

	h := New(m)

	for i, tc := range testcases {
		req := httptest.NewRequest(http.MethodGet, "https://car/id", nil)
		r := mux.SetURLVars(req, map[string]string{"id": tc.id.String()})
		w := httptest.NewRecorder()

		h.GetById(w, r)
		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(tc.resBody)) {
			t.Errorf("[Test %d] Failed Got %v Expected %v", i+1, w.Body.String(), string(tc.resBody))
		}
		if w.Code != tc.expectedStatus {
			t.Errorf("[Test %d]Failed Got %v expected %v", i+1, w.Code, tc.expectedStatus)
		}
	}
}

func TestHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockCar(ctrl)

	m.EXPECT().Get("BMW", true).Return([]models.Car{car1}, nil)

	testcases := []struct {
		param string

		resBody        []byte
		expectedStatus int
	}{
		{"?brand=BMW&IsEngine=true", []byte(`[{"id":"8f443772-132b-4ae5-9f8f-9960649b3fb4","name":"b3","year":"2003","brand":"BMW","fuel_type":"petrol","engine":{"displacement":200,"range":0,"cylinders":3,"id":"8f443772-132b-4ae5-9f8f-9960649b3fb4"}}]`), http.StatusOK},
	}

	for i, tc := range testcases {
		req := httptest.NewRequest(http.MethodPost, "https://car"+tc.param, bytes.NewReader(tc.resBody))
		w := httptest.NewRecorder()
		h := New(m)

		h.Get(w, req)

		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(tc.resBody)) {
			t.Errorf("[Test %d]Failed Got %v Expected %v", i+1, w.Body.String(), string(tc.resBody))
		}

		if w.Code != tc.expectedStatus {
			t.Errorf("[Test %d]Failed Got %v Expected %v", i+1, w.Body.String(), string(tc.resBody))
		}
	}
}

func TestHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockCar(ctrl)

	gomock.InOrder(m.EXPECT().Update(uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"), car1).Return(car1, nil))

	testcases := []struct {
		desc               string
		id                 uuid.UUID
		reqBody            []byte
		resBody            []byte
		expectedStatusCode int
	}{
		{"success", uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"), []byte(`{"car_id":"8f443772-132b-4ae5-9f8f-9960649b3fb4","name":"b3","yom":"2003","brand":"BMW","fuel_type":"petrol","engine":{"displacement":200,"range":0,"cylinders":3,"car_id":"8f443772-132b-4ae5-9f8f-9960649b3fb4"}}`), []byte(`{"car_id":"8f443772-132b-4ae5-9f8f-9960649b3fb4","name":"b3","yom":"2003","brand":"BMW","fuel_type":"petrol","engine":{"displacement":200,"range":0,"cylinders":3,"car_id":"8f443772-132b-4ae5-9f8f-9960649b3fb4"}}`), http.StatusOK},
	}

	for i, tc := range testcases {
		req := httptest.NewRequest(http.MethodPut, "https://car/id", bytes.NewReader(tc.reqBody))
		r := mux.SetURLVars(req, map[string]string{"id": tc.id.String()})
		w := httptest.NewRecorder()

		h := New(m)

		h.Update(w, r)

		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(tc.resBody)) {
			t.Errorf("[TEST%d]Failed. Got %v Expected %v", i+1, w.Body.String(), string(tc.resBody))
		}

		if w.Code != tc.expectedStatusCode {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Code, tc.expectedStatusCode)
		}
	}
}

func TestHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockCar(ctrl)

	gomock.InOrder(m.EXPECT().Create(car1).Return(car1, nil))

	testcases := []struct {
		reqBody            []byte
		resBody            []byte
		expectedStatusCode int
	}{
		{[]byte(`{"car":"8f443772-132b-4ae5-9f8f-9960649b3fb4","name":"b3","year":"2003","brand":"BMW","fuel_type":"petrol","engine":{"displacement":200,"range":0,"cylinders":3,"id":"8f443772-132b-4ae5-9f8f-9960649b3fb4"}}`), []byte(`{"id":"8f443772-132b-4ae5-9f8f-9960649b3fb4","name":"b3","year":"2003","brand":"BMW","fuel_type":"petrol","engine":{"displacement":200,"range":0,"cylinders":3,"id":"8f443772-132b-4ae5-9f8f-9960649b3fb4"}}`), http.StatusCreated},
	}

	for i, tc := range testcases {
		req := httptest.NewRequest(http.MethodPost, "https://car", bytes.NewBuffer(tc.resBody))
		w := httptest.NewRecorder()

		h := New(m)
		h.Create(w, req)

		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(tc.resBody)) {
			t.Errorf("[Test %d] Failed.Got %v\n Expected %v\n", i+1, w.Body.String(), string(tc.resBody))
		}
		if w.Code != tc.expectedStatusCode {
			t.Errorf("[Test %d] Failed.Got %v Expected %v\n", i+1, w.Code, tc.expectedStatusCode)
		}
	}
}
