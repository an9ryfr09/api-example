package app

import (
	handler "a6-api/handlers"
	"a6-api/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var app = gin.New()
	var router *gin.RouterGroup
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middleware.Configer())

	api := router.Group("/v1")
	{
		api.GET("/subjects", handler.List)
		api.GET("/subjects/:id", handler.Detail)
		api.GET("/buildings", handler.List)
		api.GET("/buildings/:id", handler.Detail)
		api.GET("/designers", handler.List)
		api.GET("/designers/:id", handler.Detail)
	}

	return app
}
