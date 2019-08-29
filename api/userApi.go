package api

import (
	"net/http"
	userDao "sample-gin-web-app/database/dao"

	"github.com/gin-gonic/gin"
)

func FindUserByID(c *gin.Context) {
	var currentUser userDao.User
	currentUser.ID = 1
	currentUser.UserName = "joe.zhang"
	currentUser.Age = 1
	currentUser.Gender = "MALE"

	c.JSON(http.StatusOK, gin.H{
		"message": currentUser,
	})
}
