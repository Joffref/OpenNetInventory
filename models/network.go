package models

import (
	"net"
	"time"

	"github.com/google/uuid"
)

type Network struct {
	ID        uuid.UUID  `json:"id"`
	NetAddr   net.IP     `json:"network_address"`
	Mask      net.IPMask `json:"mask"`
	Gateway   net.IP     `json:"gateway"`
	Protocol  string     `json:"protocol"`
	Tag       int        `json:"tag_id"`
	CreatedAt time.Time  `json:"created"`
	UpdatedAt time.Time  `json:"updated"`
}

type Networks []Network
