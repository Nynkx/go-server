package controllers

import (
	"example.com/m/v2/models"

	"github.com/gin-gonic/gin"
)

var UsersModel = new(models.Users)

type UserController struct{}

func (u UserController) GetAll(ctx *gin.Context) {

	result, err := UsersModel.GetAll()

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, result)

	return
}
