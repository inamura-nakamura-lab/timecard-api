package handler

import (
	"github.com/gin-gonic/gin"
)

type routerHandler struct {}

type IRouterHandler interface {
	SetUpRouter() *gin.Engine
}

func NewRouterHandler() IRouterHandler {
	return &routerHandler{}
}

func (r *routerHandler) SetUpRouter() *gin.Engine {
	/**
	 * Set Up Dependency Injection
	 */
	gormHandler := NewGormHandler()
	diHandler := NewDIHandler(gormHandler)
	userCtrl := diHandler.InitUserController()
	/*
	 * SetUp Routing
	 */
	router := gin.Default()
	router.GET("/user/{id}", userCtrl.GetUser)
	router.POST("/user", userCtrl.AddUser)
	router.DELETE("/user/{id}", userCtrl.DeleteUser)

	return router
}