package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/inamura-nakamura-lab/timecard-api/domain/service/interfaces"
	"log"
	"net/http"
)

type userController struct {
	interfaces.IUserService
}

type IUserController interface {
	AddUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

func NewUserController(srv interfaces.IUserService) IUserController {
	return &userController{
		srv,
	}
}

func (ctrl *userController) AddUser(ctx *gin.Context) {
	err := ctrl.IUserService.AddUser(ctx)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err})
	}
	ctx.JSON(http.StatusCreated, gin.H{})
}

func (ctrl *userController) GetUser(ctx *gin.Context) {
	result, err := ctrl.IUserService.GetUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err})
	}
	ctx.JSON(http.StatusOK, result)
}

func (ctrl *userController) DeleteUser(ctx *gin.Context) {
	err := ctrl.IUserService.DeleteUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err})
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}