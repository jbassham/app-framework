package entities

import "github.com/google/uuid"

type UUID struct {
	id uuid.UUID
}

func (u UUID) Equal(v UUID) bool {
	return u.id == v.id
}
