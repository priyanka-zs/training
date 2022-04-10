package handlers

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/types"
	"gofr_cardealership/model"
	"gofr_cardealership/services"
)

type response struct {
	Data interface{}
}
type APIHandler struct {
	serviceHandler services.Service
}

func New(car services.Service) *APIHandler {
	return &APIHandler{serviceHandler: car}
}

// Create method is used in handler layer for inserting rows.
func (h *APIHandler) Create(ctx *gofr.Context) (interface{}, error) {

	var car *model.Car
	var respCar *model.Car
	err := ctx.Bind(&car)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"Body"}}
	}

	respCar, err = h.serviceHandler.Create(ctx, car)
	if err != nil {
		return nil, err
	}
	resp := types.Response{
		Data: response{
			Data: respCar,
		},
	}

	return resp, nil
}

func (h *APIHandler) Delete(ctx *gofr.Context) (interface{}, error) {
	/*ctx := request.Context()

	writer.Header().Set("Content-Type", "application/json")

	param := mux.Vars(request)
	id := param["id"]

	uuid1, err := uuid.Parse(id) // string to uuid2 conversion.
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.Delete(ctx, uuid1)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}*/
	return nil, errors.Error("err")
}

func (h *APIHandler) Update(ctx *gofr.Context) (interface{}, error) {
	return nil, errors.Error("err")
}
func (h *APIHandler) GetByID(ctx *gofr.Context) (interface{}, error) {
	/*id := ctx.PathParam("id")

	uuid1, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	fmt.Println(uuid1, err)
	resp, err := h.serviceHandler.CarGet(ctx, uuid1)

	if err == nil {
		resp := types.Response{
			Data: response{
				Data: resp,
			},
		}
		return resp, nil
	}
	*/
	return nil, errors.Error("err")
}

func (h *APIHandler) GetByBrand(ctx *gofr.Context) (interface{}, error) {
	return nil, errors.Error("err")
}
