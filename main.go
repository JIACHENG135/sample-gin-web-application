package main

import (
	userApi "sample-gin-web-app/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:id", userApi.FindUserByID)
	r.Run() // listen and serve on 0.0.0.0:8080
}
