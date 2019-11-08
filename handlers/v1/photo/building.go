package building

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Building struct {
	Db       Photo
	building string
}

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
