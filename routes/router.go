package v1

import (
	handlerV1 "a6-api/handlers/v1"
	handlerV2 "a6-api/handlers/v2"

	"github.com/gin-gonic/gin"
)

//InitRouter initialization router group
func InitRouter(g *gin.Engine) {
	//setting router group of version 1
	v1 := g.Group("/v1")
	{
		//subject list
		v1.GET("/subjects", handlerV1.SubjectList)
		//subject detail
		v1.GET("/subject/:id", handlerV1.SubjectDetail)
		//building list
		v1.GET("/buildings", handlerV1.BuildingList)
		//building detail
		v1.GET("/building/:id", handlerV1.BuildingDetail)
		//designer list
		v1.GET("/designers", handlerV1.DesignerList)
		//designer detail
		v1.GET("/designer/:id", handlerV1.DesignerDetail)
	}

	//setting router group of version 2
	v2 := g.Group("/v2")
	{
		//subject list
		v2.GET("/subjects", handlerV2.SubjectList)
		//subject detail
		v2.GET("/subject/:id", handlerV2.SubjectDetail)
		//building list
		v2.GET("/buildings", handlerV2.BuildingList)
		//building detail
		v2.GET("/building/:id", handlerV2.BuildingDetail)
		//designer list
		v2.GET("/designers", handlerV2.DesignerList)
		//designer detail
		v2.GET("/designer/:id", handlerV2.DesignerDetail)
	}
}
