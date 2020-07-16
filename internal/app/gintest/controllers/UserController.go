package controllers

import (
	"gintest/internal/app/gintest/constants"
	"gintest/internal/app/gintest/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	BaseController
	UserService services.IUserService `inject:""`
}

func (this *UserController) List(context *gin.Context) {

	data, error := this.UserService.List()
	resData := map[string]interface{}{
		"list": data,
	}

	if error == nil {
		context.JSON(http.StatusOK, this.MakeResponse(constants.RES_CODE_SUCCESS, "", resData))
	} else {
		context.JSON(http.StatusOK, this.MakeResponse(constants.RES_CODE_OTHER, error.Error(), this.EmptyData()))
	}
}

func (this *UserController) Create(context *gin.Context) {

	name := context.PostForm("name")
	mobile := context.PostForm("mobile")
	params := map[string]string{"name": name, "mobile": mobile}

	error := this.UserService.Create(params)

	if error == nil {
		context.JSON(http.StatusOK, this.MakeResponse(constants.RES_CODE_SUCCESS, "新增成功", this.EmptyData()))
	} else {
		context.JSON(http.StatusOK, this.MakeResponse(constants.RES_CODE_OTHER, error.Error(), this.EmptyData()))
	}

}

func (this *UserController) Update(context *gin.Context) {

	idString := context.Param("id")
	id, typeError := strconv.Atoi(idString)
	if typeError != nil {
		context.JSON(http.StatusOK, this.MakeResponse(constants.RES_CODE_OTHER, typeError.Error(), this.EmptyData()))
	}

	params := map[string]string{
		"mobile": context.PostForm("mobile"),
		"name":   context.PostForm("name"),
	}
	error := this.UserService.Update(id, params)

	if error == nil {
		context.JSON(http.StatusOK, this.MakeResponse(constants.RES_CODE_SUCCESS, "更新成功", this.EmptyData()))
	} else {
		context.JSON(http.StatusOK, this.MakeResponse(constants.RES_CODE_OTHER, error.Error(), this.EmptyData()))
	}

}

func (this *UserController) Delete(context *gin.Context) {

	idString := context.Param("id")
	id, typeError := strconv.Atoi(idString)
	if typeError != nil {
		context.JSON(http.StatusOK, this.MakeResponse(constants.RES_CODE_OTHER, typeError.Error(), this.EmptyData()))
	}

	error := this.UserService.Delete(id)
	if error == nil {
		context.JSON(http.StatusOK, this.MakeResponse(constants.RES_CODE_SUCCESS, "刪除成功", this.EmptyData()))
	} else {
		context.JSON(http.StatusOK, this.MakeResponse(constants.RES_CODE_OTHER, error.Error(), this.EmptyData()))
	}
}
