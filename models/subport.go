package models

import "github.com/google/uuid"

type SubPort struct {
	ID         uuid.UUID `json:"id"`
	ParentPort Port      `json:"parent"`
	SubPort    Port      `json:"subport"`
}

type SubPorts struct {
	SubPorts []SubPort `json:"subports:"`
}
