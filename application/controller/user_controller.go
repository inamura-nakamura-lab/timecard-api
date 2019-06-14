package controller

import (
	"github.com/gin-gonic/gin"
	ictrl "github.com/inamura-nakamura-lab/timecard-api/application/controller/interfaces"
	"github.com/inamura-nakamura-lab/timecard-api/domain/service"
	"github.com/inamura-nakamura-lab/timecard-api/domain/service/interfaces"
	"github.com/inamura-nakamura-lab/timecard-api/infrastructure/persistence/repository"
	"github.com/inamura-nakamura-lab/timecard-api/utils/handler"
	"github.com/inamura-nakamura-lab/timecard-api/utils/uuid"
	"log"
	"net/http"
)

var (
	UserCtrlImpl ictrl.IUserController
)

func init() {
	log.Println("[Init]: Uuid")
	uuidImpl := uuid.NewUUIDUtil()
	log.Println("[Init]: MongoDB Connection")
	mgoConn := handler.CreateMgoConnection("user")
	log.Println("[Init]: User Controller")
	userRepoImpl := repository.NewUserRepository(mgoConn)
	userSrvImpl := service.NewUserService(userRepoImpl, uuidImpl)
	UserCtrlImpl = newUserController(userSrvImpl)
}

type userController struct {
	interfaces.IUserService
}

func newUserController(srv interfaces.IUserService) ictrl.IUserController{
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

func (ctrl *userController) GetUsers(ctx *gin.Context) {
	result, err := ctrl.IUserService.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
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
