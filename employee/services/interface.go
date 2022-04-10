package services

import (
	"context"
	"employee/models"
)

type Emp interface {
	CreateStore(ctx context.Context, emp models.Employee) (models.Employee, error)
	GetById(ctx context.Context, id int) (models.Employee, error)
}
