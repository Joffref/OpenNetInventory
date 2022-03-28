package models

import (
	"net"
	"time"

	"github.com/google/uuid"
)

type Port struct {
	ID        uuid.UUID        `json:"id"`
	Mac       net.HardwareAddr `json:"mac"`
	IP        net.IP           `json:"ip"`
	Protocol  string           `json:"protocol"`
	Tag       int              `json:"tag_id"`
	Status    string           `json:"status"`
	Tags      []string         `json:"tags"`
	CreatedAt time.Time        `json:"created"`
	UpdatedAt time.Time        `json:"updated"`
}

type Ports []Port
