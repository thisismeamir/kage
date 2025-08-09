package task_manager

import (
	"github.com/google/uuid"
)

func IdentifierGeneration(OfType string) string {
	id, _ := uuid.NewRandom()
	return OfType + "." + id.String()
}
