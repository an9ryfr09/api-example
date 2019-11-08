package router

import (
	"github.com/gin-gonic/gin"
)

//InitRouter initialization router group
func InitRouter(g *gin.Engine) {
	//setting router group of version 1
	v1 := g.Group("/v1")
	{
		v1.GET("/subjects", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "v1.test1",
			})
		})
		v1.GET("/subjects/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "v1.test2",
			})
		})
		v1.GET("/test3", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "v1.test3",
			})
		})
	}

	//setting router group of version 2
	v2 := g.Group("/v2")
	{
		v2.GET("/test1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "v2.test1",
			})
		})
		v2.GET("/test2", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "v2.test2",
			})
		})
		v2.GET("/test3", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "v2.test3",
			})
		})
	}
}
