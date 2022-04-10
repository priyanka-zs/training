package services

import (
	"context"
	"petstore/models"

	"github.com/google/uuid"
)

type Store interface {
	Post(ctx context.Context, pet models.Pet) (models.Pet, error)
	GetByID(ctx context.Context, id uuid.UUID) (models.Pet, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
