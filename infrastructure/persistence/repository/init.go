package repository

import (
	"github.com/inamura-nakamura-lab/timecard-api/domain/repository"
	"github.com/inamura-nakamura-lab/timecard-api/utils/handler"
)

func init() {
	repository.UserRepository = NewUserRepository(handler.GormConn)
}