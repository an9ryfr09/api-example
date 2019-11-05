package router

import (
	"a6-api/handler"
	"a6-api/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.RouterGroup) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.LoadConf())

	version := router.Group("/v1")
	{
		version.GET("/subjects", handler.List)
		version.GET("/subjects/:id", handler.detail)
		version.GET("/buildings", handler.List)
		version.GET("/buildings/:id", handler.detail)
		version.GET("/designers", handler.List)
		version.GET("/designers/:id", handler.detail)
	}
}
