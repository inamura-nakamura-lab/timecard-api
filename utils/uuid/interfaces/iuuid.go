package interfaces

import (
	uuid "github.com/satori/go.uuid"
)

type IUUID interface {
	GenerateUUID() uuid.UUID
}
