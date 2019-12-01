package middleware

import (
	"a6-api/utils/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = http.StatusOK
		token := c.Query("token")
		if token == "" {
			code = http.StatusBadRequest
		} else {
			claims, err := helper.ParseToken(token)
			if err != nil {
				code = http.StatusUnauthorized
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = http.StatusUnauthorized
			}
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  http.StatusText(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
