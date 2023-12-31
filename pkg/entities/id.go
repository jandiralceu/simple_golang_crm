package entities

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

// GenerateUUID generates a new UUID
func GenerateUUID() ID {
	return uuid.New()
}
