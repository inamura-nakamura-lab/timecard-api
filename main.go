package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inamura-nakamura-lab/timecard-api/application/controller"
	"github.com/inamura-nakamura-lab/timecard-api/mod"
	"log"
	"os"
)

var (
	Router *gin.Engine
)

func init() {
	log.Println("[INIT]: Gin Router")
	Router = gin.Default()
	Router.GET("/user/{id}", controller.UserCtrlImpl.GetUser)
	Router.GET("/users", controller.UserCtrlImpl.GetUsers)
	Router.POST("/user", controller.UserCtrlImpl.AddUser)
	Router.DELETE("/user/{id}", controller.UserCtrlImpl.DeleteUser)
}

func main() {
	err := Router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
