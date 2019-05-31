package controller

import "github.com/inamura-nakamura-lab/timecard-api/domain/service"

var (
	UserController IUserController
)

// Dependency Injection
func init() {
	UserController = NewUserController(service.UserService)
}