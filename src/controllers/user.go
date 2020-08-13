package controllers

import (
	"github.com/fallenstedt/gin-example/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Add(c *gin.Context) {
	user, err := models.NewUser(c, "")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		c.Abort()
		return
	}

	res, err := user.AddUser()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "id": res.InsertedID})
}

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := userModel.GetByID(c.Param("id"))
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

func (u UserController) Delete(c *gin.Context) {
	if c.Param("id") != "" {
		deleteCount, err := userModel.DeleteByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting user", "error": err})
			c.Abort()
			return
		}
		if deleteCount > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
			c.Abort()
			return
		}
		if deleteCount == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Did not delete anything!"})
			c.Abort()
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}


func (u UserController) Update( c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	user, err := models.NewUser(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		c.Abort()
		return
	}

	_, err = user.UpdateUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Bad Request"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully updated ", "user": user.ID.Hex()})
	c.Abort()
	return
}