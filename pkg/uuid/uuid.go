package uuid

import (
	"github.com/google/uuid"
)

// DefaultUUID represents the default value of an UUID field
const DefaultUUID = "11111111-1111-1111-1111-111111111111"

// MustUUID creates a new random UUID and returns it
func MustUUID() uuid.UUID {
	v, err := NewUUID()
	if err != nil {
		panic(err)
	}
	return v
}

// NewUUID creates a new random UUID and returns it with an error if exists on creation
func NewUUID() (uuid.UUID, error) {
	v, err := uuid.NewRandom()
	if err != nil {
		return uuid.UUID{}, err
	}
	return v, nil
}

// ValidUUID parses input string and returns UUID if valid, error otherwise
func ValidUUID(id string) (uuid.UUID, error) {
	validID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, err
	}
	return validID, nil
}
