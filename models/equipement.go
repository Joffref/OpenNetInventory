package models

import "github.com/google/uuid"

type Equipement struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
