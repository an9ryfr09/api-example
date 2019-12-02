package router

import (
	handler "a6-api/handlers"
	handlerV1 "a6-api/handlers/v1"
	handlerV2 "a6-api/handlers/v2"
	"a6-api/utils/loader"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//InitRouter initialization router group
func InitRouter(g *gin.Engine) {
	var protocol string
	if loader.Load().Server.EnableTLS {
		protocol = "https"
	} else {
		protocol = "http"
	}

	g.NoRoute(func(c *gin.Context) {
		desc := make([]string, 30)
		for k, r := range g.Routes() {
			desc[k] = fmt.Sprintf("%s://%s%s", protocol, loader.Load().Server.Host, r.Path)
		}
		handler.ErrorMsg(c, http.StatusBadRequest, "No matched route", desc)
	})

	//setting router group of version 1
	v1 := g.Group("/v1")
	{
		//subject list
		v1.GET("/subjects", handlerV1.SubjectList)
		//subject detail
		v1.GET("/subjects/:id", handlerV1.SubjectDetail)
		//building list
		v1.GET("/buildings", handlerV1.BuildingList)
		//building detail
		v1.GET("/buildings/:id", handlerV1.BuildingDetail)
		//designer list
		v1.GET("/designers", handlerV1.DesignerList)
		//designer detail
		v1.GET("/designers/:id", handlerV1.DesignerDetail)
		//space list
		v1.GET("/spaces", handlerV1.SpaceList)
		//style list
		v1.GET("/styles", handlerV1.StyleList)
	}

	//setting router group of version 2
	v2 := g.Group("/v2")
	{
		//subject list
		v2.GET("/subjects", handlerV2.SubjectList)
		//subject detail
		v2.GET("/subjects/:id", handlerV2.SubjectDetail)
		//building list
		v2.GET("/buildings", handlerV2.BuildingList)
		//building detail
		v2.GET("/buildings/:id", handlerV2.BuildingDetail)
		//designer list
		v2.GET("/designers", handlerV2.DesignerList)
		//designer detail
		v2.GET("/designers/:id", handlerV2.DesignerDetail)
	}
}
