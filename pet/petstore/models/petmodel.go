package models

import "github.com/google/uuid"

type Status string

const (
	Available Status = "Available"
	Pending   Status = "Pending"
	Sold      Status = "Sold"
)

type Pet struct {
	ID      uuid.UUID
	Name    string
	Age     int
	Species string
	Status  Status
}
