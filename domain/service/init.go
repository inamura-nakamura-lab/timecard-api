package service

import (
	"github.com/inamura-nakamura-lab/timecard-api/domain/repository"
	"github.com/inamura-nakamura-lab/timecard-api/domain/service/interfaces"
)

var (
	UserService interfaces.IUserService
)

func init() {
	UserService = NewUserService(repository.UserRepository)
}