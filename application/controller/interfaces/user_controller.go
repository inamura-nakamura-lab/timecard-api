package interfaces

import "github.com/gin-gonic/gin"

type IUserController interface {
	AddUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}