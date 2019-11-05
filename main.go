package main

import (
	"a6-api/middleware"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine
var RouterGroup *gin.RouterGroup

func main() {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middleware.Configer())
	//app.Run(":8080")
	// router.InitRouter(RouterGroup)
	// Router.Run(":5000")
}
