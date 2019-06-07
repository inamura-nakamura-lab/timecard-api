package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/inamura-nakamura-lab/timecard-api/domain/model"
	"github.com/inamura-nakamura-lab/timecard-api/domain/repository"
	"github.com/inamura-nakamura-lab/timecard-api/domain/service/interfaces"
	"github.com/inamura-nakamura-lab/timecard-api/infrastructure/persistence/model/mongo"
	iuuid "github.com/inamura-nakamura-lab/timecard-api/utils/uuid/interfaces"
	"log"
	"time"
)

type userService struct {
	repository.IUserRepository
	iuuid.IUUID
}

func NewUserService(repo repository.IUserRepository, uuid iuuid.IUUID) interfaces.IUserService {
	return &userService{
		repo,
		uuid,
	}
}

func (srv *userService) AddUser(ctx *gin.Context) error {
	bindUser := new(model.BindUser)
	err := ctx.Bind(&bindUser)
	log.Printf("[BindUser] %v", bindUser)
	if err != nil {
		return err
	}
	uuid := srv.IUUID.GenerateUUID()
	dbUser := &mongo.User{
		Uuid:          uuid.String(),
		Name:          bindUser.Name,
		StudentNumber: bindUser.StudentNumber,
		Date:          time.Now(),
	}
	log.Printf("[DBUser] %v", dbUser)
	return srv.IUserRepository.InsertUser(ctx, dbUser)
}

func (srv *userService) GetUsers(ctx *gin.Context) ([]*model.User, error) {
	users, err := srv.IUserRepository.SelectUsers(ctx)
	if err != nil {
		return nil, err
	}
	var result []*model.User
	for _, value := range users {
		result = append(result, &model.User{
			ID:            value.Uuid,
			Name:          value.Name,
			StudentNumber: value.StudentNumber,
			Date:          value.Date,
		})
	}
	return result, nil
}

func (srv *userService) GetUser(ctx *gin.Context) (*model.User, error) {
	paramID := ctx.Param("id")
	if paramID == "" {
		return nil, fmt.Errorf("Parameter Cannot Find")
	}
	user, err := srv.IUserRepository.SelectUser(ctx, paramID)
	if err != nil {
		return nil, err
	}
	var result *model.User
	result = &model.User{
		ID:            user.Uuid,
		Name:          user.Name,
		StudentNumber: user.StudentNumber,
		Date:          user.Date,
	}
	return result, nil
}

func (srv *userService) DeleteUser(ctx *gin.Context) error {
	paramID := ctx.Param("id")
	if paramID == "" {
		return fmt.Errorf("Parameter Cannot Find")
	}
	err := srv.IUserRepository.DeleteUser(ctx, paramID)
	if err != nil {
		return err
	}
	return nil
}
