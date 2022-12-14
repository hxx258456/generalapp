package router

import (
	"generalapp/config"
	"generalapp/internal/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/initChaincode", api.InitChaincode)
	r.Run(config.Get().System.Addr) // 监听并在 0.0.0.0:8080 上启动服务
}
