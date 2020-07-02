package entities

import (
	"time"
)

type Record interface {
	ID() int64
}

type UniversalRecord interface {
	Record

	UUID() UUID
}

type Entity interface {
	UniversalRecord

	CreatorId() int64
	CreateDate() time.Time
	ModifierId() *int64
	ModifiedDate() *time.Time
}

type EntityProvider interface {
}
