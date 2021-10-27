package main

import (
	"go-web/consumer"
	"go-web/ctrl/user"
	"go-web/producer"

	_ "go-web/producer"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	userGroup.GET("/", user.Get)
	userGroup.GET("/list", user.List)
	go producer.Start()
	go consumer.Start()
	_ = r.Run(":80")
}
