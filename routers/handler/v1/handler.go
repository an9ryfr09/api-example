package handler

import (
	model "a6-api/models"
	photo "a6-api/models/v1/photo"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var subject *photo.Subject
var building *photo.Building
var designer *photo.Designer

var modeler *model.Model

func SubjectList(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
		"data":    modeler.List(subject),
	})
}

func SubjectDetail(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
		"data":    modeler.Detail(subject),
	})
}
