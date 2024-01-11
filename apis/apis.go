package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartHttp() {
	// 设置为发布模式（初始化路由之前设置）
	gin.SetMode(gin.ReleaseMode)
	// gin 默认中间件
	r := gin.Default()

	// 访问一个错误路由时，返回404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "404, page not exists!",
		})
	})

	// 注册hello路由
	r.POST("/hello", Hello)

	// 启动API服务
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
