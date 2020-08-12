package server

import (
	"github.com/gin-gonic/gin"
	"github.com/fallenstedt/gin-example/src/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retrieve)
			userGroup.POST("/", user.Add)
		}
	}
	return router

}
