package v1

import (
	handlerV1 "a6-api/routers/handler/v1"
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
		v1.GET("/subjects/:id", handlerV1.SubjectDetail)

		// //building list
		// v1.GET("/buildings", photoV1.Building.List)
		// //building detail
		// v1.GET("/buildings/:id", photoV1.Building.Detail)

		// //designer list
		// v1.GET("/designers", photoV1.Designer.List)
		// //designer detail
		// v1.GET("/designers/:id", photoV1.Designer.Detail)
	}

	//setting router group of version 2
	// v2 := g.Group("/v2")
	// {
	// 	//subject list
	// 	v2.GET("/subjects", photoV2.Subject.List)
	// 	//subject detail
	// 	v2.GET("/subjects/:id", photoV2.Subject.Detail)

	// 	//building list
	// 	v2.GET("/buildings", photoV2.Building.List)
	// 	//building detail
	// 	v2.GET("/buildings/:id", photoV2.Building.Detail)

	// 	//designer list
	// 	v2.GET("/designers", photoV2.Designer.List)
	// 	//designer detail
	// 	v2.GET("/designers/:id", photoV2.Designer.Detail)
	// }
}
