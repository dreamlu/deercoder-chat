package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//静态目录
	r.Static("/static", "static")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run(":8007") // listen and serve on 0.0.0.0:8080
}
