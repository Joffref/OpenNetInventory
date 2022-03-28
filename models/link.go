package models

import (
	"time"

	"github.com/google/uuid"
)

type Link struct {
	ID        uuid.UUID `json:"id"`
	PortID    uuid.UUID `json:"port"`
	NetID     uuid.UUID `json:"network"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}

type Links []Link
