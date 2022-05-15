package routes

import (
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute() {
	// 设置运行模式
	gin.SetMode(utils.AppMode)
	// 初始化路由 gin.New()
	// Default 加了2个中间件
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// 设置路由组
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	// 运行
	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
