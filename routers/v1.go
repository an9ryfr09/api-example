package router

import (
	"a6-api/handler"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		v1.GET("/subjects", handler.SubjectHandler)
		v1.GET("/buildings", handler.BuildingHandler)
		v1.GET("/designers", handler.DesignerHandler)
	}
}
