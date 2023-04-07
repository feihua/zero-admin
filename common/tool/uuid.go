package tool

import "github.com/google/uuid"

func GenerateUuid() string {
	uuid := uuid.New()
	key := uuid.String()
	return key
}
