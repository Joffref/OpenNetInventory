package models

import (
	"time"

	"github.com/google/uuid"
)

type Equipement struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Ports     []Port    `json:"ports"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}

type Equipements []Equipement
