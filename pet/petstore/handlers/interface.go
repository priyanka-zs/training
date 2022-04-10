package handlers

import (
	"context"
	"petstore/models"

	"github.com/google/uuid"
)

type Service interface {
	ServicePost(ctx context.Context, pet models.Pet) (models.Pet, error)
	ServiceGetByID(ctx context.Context, id uuid.UUID) (models.Pet, error)
}
