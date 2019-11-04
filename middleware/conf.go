package middleware

import (
	"a6-api/utils"

	"github.com/gin-gonic/gin"
)

func LoadConf() gin.HandlerFunc {
	return func(c *gin.Context) {
		var conf utils.Conf
		conf.LoadConf()
	}
}
