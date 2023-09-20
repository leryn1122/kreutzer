package id

import (
	"github.com/google/uuid"
)

type Identifier interface {
	equal(ID Identifier) bool
}

type UUID struct {
	ID uuid.UUID
}

func (this_ UUID) equal(that Identifier) bool {
	return this_.ID == that.(UUID).ID
}

type Descriptor struct {
	ID string
}

func (this_ Descriptor) equal(that Identifier) bool {
	return this_.ID == that.(Descriptor).ID
}

type BigInt struct {
	ID uint64
}

func (this_ BigInt) equal(that Identifier) bool {
	return this_.ID == that.(BigInt).ID
}
