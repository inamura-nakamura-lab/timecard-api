package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/inamura-nakamura-lab/timecard-api/application/controller"
)

var (
	Router *gin.Engine
)

func init(){
	Router = gin.Default()
	Router.GET("/user/{id}", controller.UserController.GetUser)
	Router.POST("/user", controller.UserController.AddUser)
	Router.DELETE("/user/{id}", controller.UserController.DeleteUser)
}
