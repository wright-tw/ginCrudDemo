package controllers

import (
	"gintest/internal/app/gintest/services"
	// "gintest/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService services.IUserService `inject:""`
}

func (this *UserController) List(context *gin.Context) {
	data := this.UserService.List()
	context.JSON(http.StatusOK, data)
}
