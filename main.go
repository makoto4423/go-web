package main

import (
	"go-web/ctrl/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	userGroup.GET("/", user.Get)
	userGroup.GET("/list", user.List)
	_ = r.Run(":80")
}
