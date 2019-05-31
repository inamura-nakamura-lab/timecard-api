package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/inamura-nakamura-lab/timecard-api/domain/model"
	"github.com/inamura-nakamura-lab/timecard-api/domain/repository"
	"github.com/inamura-nakamura-lab/timecard-api/domain/service/interfaces"
	db "github.com/inamura-nakamura-lab/timecard-api/infrastructure/persistence/model"
	"log"
	"strconv"
)

type userService struct {
	repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) interfaces.IUserService {
	return &userService{
		repo,
	}
}

func (srv *userService) AddUser(ctx *gin.Context) error {
	bindUser := new(model.BindUser)
	err := ctx.Bind(&bindUser)
	log.Printf("[BindUser] %v", bindUser)
	if err != nil {
		return err
	}
	dbUser := &db.User{
		Name: bindUser.Name,
		StudentNum: bindUser.StudentNum,
	}
	log.Printf("[DBUser] %v", dbUser)
	return srv.IUserRepository.InsertUser(ctx, dbUser)
}

func (srv *userService) GetUser(ctx *gin.Context) (*model.User, error) {
	paramID := ctx.Param("id")
	if paramID == "" {
		return nil, fmt.Errorf("Parameter Cannot Find")
	}
	userID, err := strconv.ParseUint(paramID, 10, 32)
	if err != nil {
		return nil, err
	}
	user, err := srv.IUserRepository.SelectUser(ctx, uint(userID))
	if err != nil {
		return nil, err
	}
	result := &model.User{
		ID: user.ID,
		Name: user.Name,
		StudentNum: user.StudentNum,
		Date: user.CreatedAt,
	}
	return result, nil
}

func (srv *userService) DeleteUser(ctx *gin.Context) error {
	paramID := ctx.Param("id")
	if paramID == "" {
		return fmt.Errorf("Parameter Cannot Find")
	}
	userID, err := strconv.ParseUint(paramID, 10, 32)
	if err != nil {
		return err
	}
	err = srv.IUserRepository.DeleteUser(ctx, uint(userID))
	if err != nil {
		return err
	}
	return nil
}