package services

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/google/uuid"
	"gofr_cardealership/model"
)

type Car interface {
	CarCreate(*gofr.Context, *model.Car) (*model.Car, error)
	CarGet(*gofr.Context, uuid.UUID) (*model.Car, error)
	CarGetByBrand(*gofr.Context, string, bool) ([]*model.Car, error)
	CarDelete(*gofr.Context, uuid.UUID) error
	CarUpdate(*gofr.Context, *model.Car) (*model.Car, error)
}

type Engine interface {
	EngineCreate(*gofr.Context, *model.Engine) (*model.Engine, error)
	EngineGet(*gofr.Context, uuid.UUID) (*model.Engine, error)
	EngineDelete(*gofr.Context, uuid.UUID) error
	EngineUpdate(*gofr.Context, *model.Engine) (*model.Engine, error)
}
