package models

import (
	"net"
	"time"

	"github.com/google/uuid"
)

type Link struct {
	ID        uuid.UUID  `json:"id"`
	Interface *Port      `json:"interface"`
	Network   net.IP     `json:"network"`
	Mask      net.IPMask `json:"mask"`
	CreatedAt time.Time  `json:"created"`
	UpdatedAt time.Time  `json:"updated"`
}

type Links struct {
	Links []Link `json:"links"`
}

type LinksCluster struct {
	ID        uuid.UUID `json:"id"`
	Topic     string    `json:"topic"`
	Links     *Links    `json:"links"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}
