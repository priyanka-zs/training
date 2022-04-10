package handlers

import (
	"context"

	"github.com/google/uuid"

	"github.com/zopsmart/GoLang-Interns-2022/model"
)

type Car interface {
	Create(context.Context, *model.Car) (*model.Car, error)
	GetByID(context.Context, uuid.UUID) (*model.Car, error)
	Get(context.Context, string, bool) ([]*model.Car, error)
	Delete(context.Context, uuid.UUID) error
	Update(context.Context, *model.Car) (*model.Car, error)
}
