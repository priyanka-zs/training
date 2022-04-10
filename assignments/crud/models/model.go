package models

import "github.com/google/uuid"

type Brand string

const (
	tesla    Brand = "telsa"
	porsche        = "porsche"
	ferrari        = "ferrari"
	mercedes       = "mercedes"
	bmw            = "bmw"
)

type fuelType string

const (
	electric fuelType = "electric"
	petrol            = "petrol"
	diesel            = "diesel"
)

type Engine struct {
	Displacement  float64
	NoOfCylinders int
	Ranges        float64
}
type Car struct {
	Id     uuid.UUID
	Name   string
	Year   int
	Brand  Brand
	Fuel   fuelType
	Engine Engine
}
