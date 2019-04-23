package handler

import (
	"github.com/inamura-nakamura-lab/timecard-api/application/controller"
	"github.com/inamura-nakamura-lab/timecard-api/domain/service"
	"github.com/inamura-nakamura-lab/timecard-api/infrastructure/persistence/repository"
)

type diHandler struct {
	IGormHandler
}

type IDIHandler interface {
	InitUserController() controller.IUserController
}

func NewDIHandler(orm IGormHandler) IDIHandler {
	return &diHandler{
		orm,
	}
}

func (di *diHandler) InitUserController() controller.IUserController {
	db, err := di.CreateConnection()
	if err != nil {
		panic(err)
	}
	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(repo)
	return controller.NewUserController(srv)
}