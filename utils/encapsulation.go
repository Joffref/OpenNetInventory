package utils

import (
	"errors"

	"github.com/joffref/netrman/models"
)

func (p models.Port) CheckVlanNumber() error {
	if p.TagID.ID < 0 || p.TagID.ID > 4095 {
		return errors.New("Wrong VLAN number")
	}
	return nil
}

func (p models.Port) CheckVXlanNumber() error {
	if p.TagID.ID < 0 || p.TagID.ID > 16777215 {
		return errors.New("Wrong VXLAN number")
	}
	return nil
}
