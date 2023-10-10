package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func route() {
	// 创建一个服务
	ginServer := gin.Default()
	// 网页图标
	ginServer.Use(favicon.New("./test.ico"))

	// 连接数据库的代码

	// 路由
	ginServer.GET("/test", func(context *gin.Context) {
		// 重定向 301
		context.Redirect(http.StatusMovedPermanently, "https://cn.bing.com/")
	})
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	// 路由组
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("add")
		userGroup.POST("login")
		userGroup.POST("logout")
	}
	orderGroup := ginServer.Group("/order")
	{
		orderGroup.GET("/add")
		orderGroup.POST("/delete")
	}

	// 服务器端口
	ginServer.Run(":8082")
}
