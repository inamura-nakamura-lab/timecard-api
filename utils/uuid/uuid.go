package uuid

import (
	"github.com/inamura-nakamura-lab/timecard-api/utils/uuid/interfaces"
	uuid "github.com/satori/go.uuid"
)


type uuidUtil struct{}

func NewUUIDUtil() interfaces.IUUID {
	return &uuidUtil{}
}

func (uid *uuidUtil) GenerateUUID() uuid.UUID {
	a := uuid.NewV4()
	u1 := uuid.Must(a, nil)
	return u1
}
