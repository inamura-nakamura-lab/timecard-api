package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/inamura-nakamura-lab/timecard-api/domain/model"
)

type IUserService interface {
	AddUser(ctx *gin.Context) error
	GetUsers(ctx *gin.Context) ([]*model.User, error)
	GetUser(ctx *gin.Context) (*model.User, error)
	DeleteUser(ctx *gin.Context) error
}
