package services

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/google/uuid"
	"gofr_cardealership/model"
)

type Service struct {
	engine Engine
	car    Car
}

func New(c Car, e Engine) *Service {
	return &Service{car: c, engine: e}
}

// Validation function is used to validate the car details
func Validation(car model.Car) error { //nolint
	if car.Year < 1980 || car.Year > 2022 {
		return errors.Error("year should be between 1980 and 2022")
	}

	if car.Brand != model.Tesla && car.Brand != model.Porsche && car.Brand != model.Ferrari &&
		car.Brand != model.Mercedes && car.Brand != model.Bmw {
		return errors.Error("invalid brand")
	}

	if car.Fuel != model.Electric && car.Fuel != model.Petrol && car.Fuel != model.Diesel {
		return errors.Error("invalid fuel")
	}

	if car.Fuel == model.Petrol || car.Fuel == model.Diesel {
		if car.Engine.Displacement <= 0 {
			return errors.Error("displacement should be positive")
		}

		if car.Engine.NoOfCylinders <= 0 {
			return errors.Error("noOfCylinders should be positive")
		}

		if car.Engine.Ranges != 0 {
			return errors.Error("given fuel type does not support range")
		}
	}

	if car.Fuel == "electric" {
		if car.Engine.Displacement != 0 {
			return errors.Error("electric cannot have displacement")
		}

		if car.Engine.NoOfCylinders != 0 {
			return errors.Error("electric cannot have  noOfCylinders")
		}

		if car.Engine.Ranges < 0 {
			return errors.Error("range should be positive")
		}
	}

	return nil
}

// Create validates the input and creates a row
func (s *Service) Create(ctx *gofr.Context, car *model.Car) (*model.Car, error) { //nolint
	err := Validation(*car)
	if err != nil {
		return &model.Car{}, err
	}

	result, err := s.car.CarCreate(ctx, car)

	if err != nil {
		return &model.Car{}, err
	}

	car.Engine.EngineID = result.Engine.EngineID

	resEngine, err := s.engine.EngineCreate(ctx, &car.Engine)

	if err != nil {
		return &model.Car{}, err
	}

	result.Engine = *resEngine

	return result, nil
}

func (s *Service) GetByID(ctx *gofr.Context, id uuid.UUID) (*model.Car, error) {

	r, err := s.car.CarGet(ctx, id)

	if err != nil {
		return &model.Car{}, err
	}

	eng, err := s.engine.EngineGet(ctx, r.Engine.EngineID)

	if err != nil {
		return &model.Car{}, err
	}

	r.Engine = *eng

	return r, nil
}

// Delete takes id as input and deletes the corresponding row
func (s *Service) Delete(ctx *gofr.Context, id uuid.UUID) error {
	car, err := s.car.CarGet(ctx, id)
	if err != nil {
		return err
	}

	err = s.engine.EngineDelete(ctx, car.Engine.EngineID)
	if err != nil {
		return errors.Error("invalid id")
	}

	err = s.car.CarDelete(ctx, id)
	if err != nil {
		return errors.Error("invalid id")
	}

	return nil
}

// Update method is used to update a row in the database
func (s *Service) Update(ctx *gofr.Context, car *model.Car) (*model.Car, error) { //nolint
	err := Validation(*car)
	if err != nil {
		return &model.Car{}, err
	}

	resp, err := s.car.CarUpdate(ctx, car)
	if err != nil {
		return &model.Car{}, err
	}

	res, err := s.car.CarGet(ctx, car.ID)
	if err != nil {
		return &model.Car{}, err
	}

	car.Engine.EngineID = res.Engine.EngineID
	resEngine, err := s.engine.EngineUpdate(ctx, &car.Engine)

	if err != nil {
		return &model.Car{}, err
	}

	resp.Engine = *resEngine

	return resp, nil
}

//	Get takes brand as input and returns the corresponding row
func (s Service) Get(ctx *gofr.Context, brand string, isEngine bool) ([]*model.Car, error) { //nolint
	if brand != "tesla" && brand != "porsche" && brand != "ferrari" &&
		brand != "mercedes" && brand != "bmw" {
		return []*model.Car{}, errors.Error("invalid brand")
	}

	res, err := s.car.CarGetByBrand(ctx, brand, isEngine)
	if err != nil {
		return []*model.Car{}, err
	}

	if isEngine {
		for i := 0; i < len(res); i++ {
			engine, er := s.engine.EngineGet(ctx, res[i].Engine.EngineID)

			if er != nil {
				return []*model.Car{}, er
			}

			res[i].Engine = *engine
		}
	}

	return res, nil
}
