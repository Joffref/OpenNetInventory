package models

import (
	"time"

	"github.com/google/uuid"
)

type Equipement struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Constructor  string    `json:"vendor"`
	Version      string    `json:"version"`
	Ports        Ports     `json:"ports"`
	Links        Links     `json:"links"`
	SubscribedAt time.Time `json:"subscribed"`
	UpdatedAt    time.Time `json:"updated"`
}
