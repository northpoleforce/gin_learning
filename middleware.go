package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"log"
)

// 自定义中间件 拦截器
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 通过自定义的中间件，设置的值，在后续处理只要调用了这个中间件都可以拿到这里的参数
		context.Set("usersession", "userid-1")
		context.Next() // 放行
		//context.Abort() // 阻止
	}
}

func middleware() {
	// 创建一个服务
	ginServer := gin.Default()
	// 网页图标
	ginServer.Use(favicon.New("./test.ico"))

	// 连接数据库的代码

	ginServer.GET("/hello", myHandler(), func(context *gin.Context) {
		usersession := context.MustGet("usersession").(string)
		log.Println("===========>", usersession)
		context.JSON(200, gin.H{"msg": "hello,world"})
	})

	// 服务器端口
	ginServer.Run(":8082")
}
