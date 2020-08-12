package controllers

import (
	"github.com/fallenstedt/gin-example/src/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Add(c *gin.Context) {
	var us models.User
	err := c.BindJSON(&us)
	if err != nil {
		log.Printf("Unable to parse body %f", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		c.Abort()
		return
	}
	log.Printf("Got User %v", us.Name)
	c.JSON(http.StatusNoContent, gin.H{})
}

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := userModel.GetByID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User found!", "user": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}
