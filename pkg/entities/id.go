package entities

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

// GenerateUUID generates a new UUID
func GenerateUUID() ID {
	return uuid.New()
}

// ParseUUID parses a string into a UUID
//func ParseUUID(s string) (ID, error) {
//	id, err := uuid.Parse(s)
//	return id, err
//}
