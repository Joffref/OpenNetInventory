package models

import (
	"time"

	"github.com/google/uuid"
)

type Cluster struct {
	ID        uuid.UUID                `json:"id"`
	Type      string                   `json:"type"`
	Topic     string                   `json:"topic"`
	Objects   []map[string]interface{} `json:"objects"`
	CreatedAt time.Time                `json:"created"`
	UpdatedAt time.Time                `json:"updated"`
}

type Clusters []Cluster
