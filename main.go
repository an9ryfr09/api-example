package main

import (
	"a6-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(middleware.LoadConf())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
