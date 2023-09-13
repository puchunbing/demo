package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个Gin路由器
	r := gin.Default()

	// 定义路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	r.POST("/submit", func(c *gin.Context) {
		var json struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		if err := c.Bind(&json); err == nil {
			c.JSON(200, gin.H{
				"message": "JSON received",
				"name":    json.Name,
				"email":   json.Email,
			})
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
		}
	})

	// 启动Gin应用程序
	r.Run()
}
