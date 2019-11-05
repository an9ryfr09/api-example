package middleware

import (
	"a6-api/conf"

	"github.com/gin-gonic/gin"
)

func Configer() gin.HandlerFunc {
	var conf conf.AppConf
	conf.Load()
	return func(c *gin.Context) {
		c.Next()
	}
}
