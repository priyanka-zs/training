package handlers

import (
	"bytes"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gofr_cardealership/model"
	"gofr_cardealership/services"
	_ "gofr_cardealership/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHandler_Create is used to test Create in handler
func TestHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := services.NewMockCar()
	h := New(m)
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
		{"invalid id", &car2, &model.Car{}, errors.Error("error from service layer"), http.StatusBadRequest},
	}

	for _, tc := range testcases {
		body, err := json.Marshal(tc.reqBody)
		if err != nil {
			return
		}

		req := httptest.NewRequest(http.MethodPost, "https://car", bytes.NewBuffer(body))

		m.EXPECT().Create(gomock.Any(), tc.reqBody).Return(tc.resBody, tc.err)
		ctx := gofr.NewContext(responder.NewContextualResponder(httptest.NewRecorder(), req), request.NewHTTPRequest(req), nil)
		_, err = h.Create(ctx)
		if err != nil {
			return
		}
		assert.Equal(t, tc.err, err)
	}
}

// TestMockService_CreateError is used to test Create in handler
//func TestMockService_CreateError(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//	m := services.NewMockCar(ctrl)
//	testcases := []struct {
//		desc   string
//		input  []byte
//		status int
//		err    error
//	}{
//		{"missing parameters", []byte(`[{"name": "bmw 2", "year": 2000}]`), http.StatusBadRequest, errors.InvalidParam{Param: []string{"Body"}}},
//	}
//
//	for _, tc := range testcases {
//		req := httptest.NewRequest(http.MethodPost, "https://car", bytes.NewBuffer(tc.input))
//		ctx := gofr.NewContext(responder.NewContextualResponder(httptest.NewRecorder(), req), request.NewHTTPRequest(req), nil)
//		h := New(m)
//		_, err := h.Create(ctx)
//		if err != nil {
//			return
//		}
//
//		assert.Equal(t, tc.err, err)
//	}
//}
//
//// TestHandler_GetById is used to test GetById method in handler layer.
//func TestHandler_GetByID(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//	m := services.NewMockCar(ctrl)
//	h := New(m)
//	ID := uuid.New()
//	ID1 := uuid.New()
//	var car1 = model.Car{
//		ID: ID, Name: "bmw3", Year: 2005, Brand: "bmw",
//		Fuel: "petrol", Engine: model.Engine{Displacement: 150, Ranges: 0, NoOfCylinders: 3,
//			EngineID: ID1}}
//
//	testcases := []struct {
//		desc           string
//		id             uuid.UUID
//		resBody        *model.Car
//		err            error
//		expectedStatus int
//	}{
//		{"success case", ID, &car1, nil, http.StatusOK},
//		//{"id not given", uuid.Nil, &model.Car{}, errors.Error("error"), http.StatusBadRequest},
//	}
//
//	for _, tc := range testcases {
//		req := httptest.NewRequest(http.MethodGet, "https://car/id", nil)
//		ctx := gofr.NewContext(responder.NewContextualResponder(httptest.NewRecorder(), req), request.NewHTTPRequest(req), nil)
//		ctx.SetPathParams(map[string]string{"id": tc.id.String()})
//		m.EXPECT().CarGet(gomock.Any(), tc.id).Return(tc.resBody, tc.err)
//		_, err := h.GetByID(ctx)
//		if err != nil {
//			return
//		}
//
//		assert.Equal(t, tc.err, err)
//	}
//}
//
//// TestHandler_GetByIDErr is used to test GetByID
//func TestHandler_GetByIDErr(t *testing.T) {
//	ctrl := gomock.NewController(t)
//
//	defer ctrl.Finish()
//	m := services.NewMockCar(ctrl)
//	h := New(m)
//
//	testcases := []struct {
//		desc           string
//		id             string
//		resBody        *model.Car
//		err            error
//		expectedStatus int
//	}{
//		{"id not given", "123", &model.Car{}, errors.Error("invalid uuid"), http.StatusBadRequest},
//	}
//
//	for _, tc := range testcases {
//		req := httptest.NewRequest(http.MethodGet, "https://car/id", nil)
//		ctx := gofr.NewContext(responder.NewContextualResponder(httptest.NewRecorder(), req), request.NewHTTPRequest(req), nil)
//		ctx.SetPathParams(map[string]string{"id": tc.id})
//		_, err := h.GetByID(ctx)
//		if err != nil {
//			return
//		}
//
//		assert.Equal(t, tc.err, err)
//	}
//}
//
//// TestHandler_Delete is used to test Delete method in handler layer.
//func TestHandler_Delete(t *testing.T) {
//
//	ID := uuid.New()
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	m := services.NewMockCar(ctrl)
//	h := New(m)
//	testcases := []struct {
//		desc     string
//		id       uuid.UUID
//		expected int
//		err      error
//	}{
//		{"success case", ID, http.StatusNoContent, nil},
//		{"id not given", uuid.Nil, http.StatusBadRequest, errors.New("error")},
//	}
//
//	for _, tc := range testcases {
//		req := httptest.NewRequest(http.MethodGet, "https://car", nil)
//		ctx := gofr.NewContext(responder.NewContextualResponder(httptest.NewRecorder(), req), request.NewHTTPRequest(req), nil)
//		ctx.SetPathParams(map[string]string{"id": tc.id.String()})
//
//		gomock.InOrder(m.EXPECT().CarDelete(gomock.Any(), tc.id).Return(tc.err))
//		_, err := h.Delete(ctx)
//		assert.Equal(t, tc.err, err)
//
//	}
//}

//// TestHandler_DeleteError is used to test delete error case
//func TestHandler_DeleteError(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	m := NewMockCar(ctrl)
//
//	testcases := []struct {
//		desc   string
//		id     string
//		status int
//	}{
//		{"id invalid", "123", http.StatusBadRequest},
//	}
//
//	for _, tc := range testcases {
//		req := httptest.NewRequest(http.MethodDelete, "https://car/id", nil)
//		w := httptest.NewRecorder()
//
//		h := New(m)
//		h.Delete(w, req)
//
//		assert.Equal(t, tc.status, w.Code)
//	}
//}
