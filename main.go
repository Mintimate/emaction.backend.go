package main

import (
	"emaction/internal/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// 设置 CORS 跨域
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, PATCH, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// 设置 /reactions 开头路由
	reactionGroup := router.Group("/")
	{
		reactionGroup.GET("/reactions", controller.GetReactions)
		reactionGroup.PATCH("/reaction", controller.UpdateReaction)
	}

	// 启动服务器
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
