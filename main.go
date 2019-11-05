package main

import (
	"a6-api/middleware"
	router "a6-api/router"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine
var RouterGroup *gin.RouterGroup

func main() {
	router.Router(RouterGroup)
	middleware.LoadConf()
	Router.Run(":5000")
}

func loadRouter() {
	apiv2 := Router.Group("/api/v1")
	rest := apiv2.Group("/rest")
	router.Router(rest)

}
