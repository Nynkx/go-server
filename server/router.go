package server

import (
	"example.com/m/v2/controllers"
	"github.com/gin-gonic/gin"
)

func addUserRoutes(router *gin.Engine){
	user := new(controllers.UserController)

	router.GET("/test", user.GetAll)
}

func NewRouter() *gin.Engine{
	router := gin.New()

	router.Use(gin.Logger())

	addUserRoutes(router)

	

	return router
}