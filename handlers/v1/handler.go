package handler

import (
	model "a6-api/models"
	photo "a6-api/models/v1/photo"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var modeler *model.Model
var subject *photo.Subject

var SubjectList = func(c *gin.Context) {
	data := modeler.List(subject)
	defer c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
		"data":    data,
	})
}

var SubjectDetail = func(c *gin.Context) {
	data := modeler.Detail(subject)
	defer c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
		"data":    data,
	})
}
