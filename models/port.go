package models

import (
	"errors"
	"net"
	"time"

	"github.com/google/uuid"
)

type TagProto string

const (
	VLAN  TagProto = "vlan"
	VXLAN TagProto = "vxlan"
)

type Port struct {
	ID       uuid.UUID        `json:"id"`
	Name     string           `json:"name"`
	State    string           `json:"state"`
	IPaddr   net.IP           `json:"ip"`
	MAC      net.HardwareAddr `json:"mac"`
	Tags     int              `json:"tags"`
	TagProto TagProto         `json:"encapsulation"`
	TagID    struct {
		ID int `json:"id"`
	}
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}

type Ports struct {
	PortsList []Port `json:"ports"`
}

type PortsCluster struct {
	ID        uuid.UUID `json:"id"`
	Topic     string    `json:"topic"`
	Ports     *Ports    `json:"ports"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}

func (p Port) CheckVlanNumber() error {
	if p.TagID.ID < 0 || p.TagID.ID > 4095 {
		return errors.New("wrong vlan number")
	}
	return nil
}

func (p Port) CheckVXlanNumber() error {
	if p.TagID.ID < 0 || p.TagID.ID > 16777215 {
		return errors.New("wrong vxlan number")
	}
	return nil
}
