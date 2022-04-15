package models

import (
	"net"
)

type Network struct {
	ID        string `json:"id"`
	NetAddr   net.IP `json:"network_address"`
	Mask      net.IP `json:"mask"`
	Gateway   net.IP `json:"gateway"`
	Protocol  string `json:"protocol"`
	Tag       int    `json:"tag_id"`
	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`
}
