package model

import "github.com/google/uuid"

type Engine struct {
	EngineID      uuid.UUID `json:"EngineID"`
	Displacement  float64   `json:"Displacement,omitempty"`
	NoOfCylinders int       `json:"NoOfCylinders,omitempty"`
	Ranges        float64   `json:"Ranges,omitempty"`
}
