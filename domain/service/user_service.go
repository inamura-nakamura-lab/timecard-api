package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/inamura-nakamura-lab/timecard-api/domain/model"
	"github.com/inamura-nakamura-lab/timecard-api/domain/repository"
	db "github.com/inamura-nakamura-lab/timecard-api/infrastructure/persistence/model"
	"strconv"
)

type userService struct {
	repository.IUserRepository
}

type IUserService interface {
	AddUser(ctx *gin.Context) error
	GetUser(ctx *gin.Context) (*model.User, error)
	DeleteUser(ctx *gin.Context) error
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &userService{
		repo,
	}
}

func (srv *userService) AddUser(ctx *gin.Context) error {
	bindUser := new(model.BindUser)
	err := ctx.Bind(&bindUser)
	if err != nil {
		return err
	}
	dbUser := &db.User{
		Name: bindUser.Name,
		StudentNum: bindUser.Name,
	}
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