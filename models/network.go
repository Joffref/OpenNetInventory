package models

import (
	"net"
	"time"
)

type Network struct {
	ID        string    `json:"id" binding:"required,uuid"`
	NetAddr   net.IP    `json:"network_address"`
	Mask      net.IP    `json:"mask"`
	Gateway   net.IP    `json:"gateway"`
	Protocol  string    `json:"protocol"`
	Tag       int       `json:"tag_id"`
	CreatedAt time.Time `json:"created" time_format:"2006-01-02T15:04:05Z"`
	UpdatedAt time.Time `json:"updated" time_format:"2006-01-02T15:04:05Z"`
}
