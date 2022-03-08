package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (u UserController) GetAll(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"a":"a"})
	return
}