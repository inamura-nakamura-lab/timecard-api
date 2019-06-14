package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/inamura-nakamura-lab/timecard-api/domain/model/request"
	"github.com/inamura-nakamura-lab/timecard-api/infrastructure/persistence/model/mongo"
)

type IUserRepository interface {
	InsertUser(ctx *gin.Context, user *mongo.User) error
	SelectUsers(ctx *gin.Context) ([]*mongo.User, error)
	SelectUser(ctx *gin.Context, userID string) (*mongo.User, error)
	DeleteUser(ctx *gin.Context, userID string) error
	SelectAttendance(ctx *gin.Context, userID string) (*mongo.OrgAttendance, error)
	InsertAttendance(ctx *gin.Context, attendance *request.PostAttendance) error
}
