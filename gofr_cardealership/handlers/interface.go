package handlers

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/google/uuid"
	"gofr_cardealership/model"
)

type Car interface {
	Create(*gofr.Context, *model.Car) (*model.Car, error)
	GetByID(*gofr.Context, uuid.UUID) (*model.Car, error)
	Get(*gofr.Context, string, bool) ([]*model.Car, error)
	Delete(*gofr.Context, uuid.UUID) error
	Update(*gofr.Context, *model.Car) (*model.Car, error)
}
