package main

import (
	"ginblog/model"
	"ginblog/routes"
)

// 项目主入口

func main() {
	// 初始化数据库
	model.InitDb()

	// 初始化路由
	routes.InitRoute()
}
