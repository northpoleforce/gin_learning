package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func get_params() {
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

	// 接受前端传递过来的参数
	// url?userid=007&username=northpoleforce
	ginServer.GET("/user/info", func(context *gin.Context) {
		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})
	// /user/info/007/northpoleforce
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// 掌握技术后面的应用 - 掌握基础知识，加以了解web开发
	// 前端给后端传递 json
	ginServer.POST("/json", func(context *gin.Context) {
		// request.body
		// []byte
		data, _ := context.GetRawData()
		var m map[string]interface{}
		// 包装为json数据 []byte
		_ = json.Unmarshal(data, &m)
		context.JSON(http.StatusOK, m)
	})

	// 支持函数式编程 =>
	ginServer.POST("/user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	// 服务器端口
	ginServer.Run(":8082")
}
