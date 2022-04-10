package model

import "github.com/google/uuid"

type Brand string

const (
	Tesla    Brand = "tesla"
	Porsche  Brand = "porsche"
	Ferrari  Brand = "ferrari"
	Mercedes Brand = "mercedes"
	Bmw      Brand = "bmw"
)

type fuelType string

const (
	Electric fuelType = "electric"
	Petrol   fuelType = "petrol"
	Diesel   fuelType = "diesel"
)

type Car struct {
	ID     uuid.UUID `json:"ID"`
	Name   string    `json:"Name"`
	Year   int       `json:"Year"`
	Brand  Brand     `json:"Brand"`
	Fuel   fuelType  `json:"Fuel"`
	Engine Engine    `json:"Engine"`
}
