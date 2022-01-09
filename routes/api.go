package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		// 注册测试路由
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "world",
			})
		})
	}
}
