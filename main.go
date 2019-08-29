package main

import (
	userApi "sample-gin-web-app/api"
	db "sample-gin-web-app/database"
	"sample-gin-web-app/database/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.GetDBConnection()

	if err != nil {
		panic("fatal error, failed to connect to database")
	}
	db.AutoMigrate(&dao.User{})
	defer db.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:id", userApi.FindUserByID)
	r.Run() // listen and serve on 0.0.0.0:8080
}
