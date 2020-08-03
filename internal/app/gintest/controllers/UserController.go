package controllers

import (
	"gintest/internal/app/gintest/constants"
	"gintest/internal/app/gintest/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewUserController(service *services.UserService) *UserController {
	return &UserController{UserService: service}
}

type UserController struct {
	BaseController
	UserService services.IUserService
}

func (u *UserController) List(context *gin.Context) {
	data, error := u.UserService.List()
	resData := map[string]interface{}{
		"list": data,
	}

	if error == nil {
		context.JSON(http.StatusOK, u.MakeResponse(constants.ResCodeSuccess, "", resData))
	} else {
		context.JSON(http.StatusOK, u.MakeResponse(constants.ResCodeOther, error.Error(), u.EmptyData()))
	}
}

func (u *UserController) Create(context *gin.Context) {
	name := context.PostForm("name")
	mobile := context.PostForm("mobile")
	params := map[string]string{"name": name, "mobile": mobile}

	error := u.UserService.Create(params)

	if error == nil {
		context.JSON(http.StatusOK, u.MakeResponse(constants.ResCodeSuccess, "新增成功", u.EmptyData()))
	} else {
		context.JSON(http.StatusOK, u.MakeResponse(constants.ResCodeOther, error.Error(), u.EmptyData()))
	}
}

func (u *UserController) Update(context *gin.Context) {
	idString := context.Param("id")
	id, typeError := strconv.Atoi(idString)
	if typeError != nil {
		context.JSON(http.StatusOK, u.MakeResponse(constants.ResCodeOther, typeError.Error(), u.EmptyData()))
	}

	params := map[string]string{
		"mobile": context.PostForm("mobile"),
		"name":   context.PostForm("name"),
	}
	error := u.UserService.Update(id, params)

	if error == nil {
		context.JSON(http.StatusOK, u.MakeResponse(constants.ResCodeSuccess, "更新成功", u.EmptyData()))
	} else {
		context.JSON(http.StatusOK, u.MakeResponse(constants.ResCodeOther, error.Error(), u.EmptyData()))
	}
}

func (u *UserController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, typeError := strconv.Atoi(idString)
	if typeError != nil {
		context.JSON(http.StatusOK, u.MakeResponse(constants.ResCodeOther, typeError.Error(), u.EmptyData()))
	}

	error := u.UserService.Delete(id)
	if error == nil {
		context.JSON(http.StatusOK, u.MakeResponse(constants.ResCodeSuccess, "刪除成功", u.EmptyData()))
	} else {
		context.JSON(http.StatusOK, u.MakeResponse(constants.ResCodeOther, error.Error(), u.EmptyData()))
	}
}
