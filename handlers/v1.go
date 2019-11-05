package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SubjectHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
		"data":{}
	})
}

func BuildingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
		"data":{}
	})
}

func DesignerHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
		"data":{}
	})
}
