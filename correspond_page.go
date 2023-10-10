package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func correspond_page() {
	// 创建一个服务
	ginServer := gin.Default()
	// 网页图标
	ginServer.Use(favicon.New("./test.ico"))

	// 连接数据库的代码

	// 加载静态页面
	ginServer.LoadHTMLGlob("templates/*")
	// 另外一种加载方法
	//ginServer.LoadHTMLFiles("templates/index.html")
	// 加载资源文件
	ginServer.Static("/static", "./static")

	// 响应一个页面给前端
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "这是go后台传递来的数据",
		})
	})

	// 服务器端口
	ginServer.Run(":8082")
}
