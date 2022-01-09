package main

import (
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	bootstrap.SetRoute(r)

	r.Run(":8888")
}
