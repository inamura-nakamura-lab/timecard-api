package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/inamura-nakamura-lab/timecard-api/infrastructure/persistence/model"
)

type IUserRepository interface {
	InsertUser(ctx *gin.Context, user *model.User) error
	SelectUser(ctx *gin.Context, userID uint) (*model.User, error)
	DeleteUser(ctx *gin.Context, userID uint) error
}