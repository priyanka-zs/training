package services

import (
	"context"
	"github.com/zopsmart/GoLang-Interns-2022/model"

	"github.com/google/uuid"
)

type Car interface {
	CarCreate(context.Context, *model.Car) (*model.Car, error)
	CarGet(context.Context, uuid.UUID) (*model.Car, error)
	CarGetByBrand(context.Context, string, bool) ([]*model.Car, error)
	CarDelete(context.Context, uuid.UUID) error
	CarUpdate(context.Context, *model.Car) (*model.Car, error)
}

type Engine interface {
	EngineCreate(context.Context, *model.Engine) (*model.Engine, error)
	EngineGet(context.Context, uuid.UUID) (*model.Engine, error)
	EngineDelete(context.Context, uuid.UUID) error
	EngineUpdate(context.Context, *model.Engine) (*model.Engine, error)
}
