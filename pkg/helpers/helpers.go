package helpers

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	guid := uuid.New()
	return guid.String()
}
