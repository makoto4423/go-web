package main

import (
	// "go-web/consumer"
	"go-web/ctrl/user"
	"go-web/filter"
	"go-web/redis"
	"go-web/zookeeper"

	// "go-web/producer"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(filter.Filter())
	userGroup := r.Group("/user")
	userGroup.GET("/", user.Get)
	userGroup.GET("/list", user.List)
	// go producer.Start()
	// go consumer.Start()
	redis.Get()
	zookeeper.Get()
	_ = r.Run(":80")
}
