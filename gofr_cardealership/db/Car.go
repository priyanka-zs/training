package db

import (
	"github.com/google/uuid"
	"gofr_cardealership/model"
)

var Car = []model.Car{
	{ID: uuid.New(), Name: "tesla 1", Year: 2020, Brand: "tesla", Fuel: "diesel"},
}
