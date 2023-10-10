// package learning
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

//func hello() {
//	fmt.Println("Hello World!")
//}

func web_hello() {
	// 创建一个服务
	ginServer := gin.Default()
	// 网页图标
	ginServer.Use(favicon.New("./test.ico"))

	// 连接数据库的代码

	// 访问地址，处理请求 Request Response
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello,world"})
	})
	// 服务器端口
	ginServer.Run(":8082")
}
