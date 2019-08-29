package api

import (
	"net/http"
	userDao "sample-gin-web-app/database/dao"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindUserByID(c *gin.Context) {
	var currentUser userDao.User
	userID, _ := strconv.Atoi(c.Param("id"))
	currentUser = userDao.GetUserByID(userID)

	c.JSON(http.StatusOK, gin.H{
		"message": currentUser,
	})
}
