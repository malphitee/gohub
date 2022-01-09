package bootstrap

import (
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 路由初始化
func SetRoute(router *gin.Engine) {
	registerGlobalMiddleware(router)
	routes.RegisterApiRoutes(router)
	set404Handler(router)
}

// 注册公共的中间件
func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func set404Handler(router *gin.Engine) {
	// 处理404请求
	router.NoRoute(func(c *gin.Context) {
		// 获取表头的Accept信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "未找到页面")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_coce":    http.StatusNotFound,
				"error_message": "路由未定义，请确认url和请求方法是否正确",
			})
		}
	})
}
