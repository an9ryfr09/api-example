package main

import (
	"a6-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middleware.Configer())
	// //app.Run(":8080")
	// app.InitRouter(routerGroup)
}
