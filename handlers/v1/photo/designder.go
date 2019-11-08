package designer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
	})
}

func Detail(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
	})
}
