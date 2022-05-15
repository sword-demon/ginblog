package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRoute() {
	// 设置运行模式
	gin.SetMode(utils.AppMode)
	// 初始化路由 gin.New()
	// Default 加了2个中间件
	r := gin.Default()

	v1Group := r.Group("api/v1")
	{
		// 用户模块的路由接口
		v1Group.POST("/user/add", v1.AddUser)
		v1Group.GET("/users", v1.GetUsers)
		v1Group.PUT("/user/:id", v1.EditUser)
		v1Group.DELETE("/user/:id", v1.DeleteUser)
		// 分类模块的路由接口

		// 文章模块的路由接口
	}

	// 运行
	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
