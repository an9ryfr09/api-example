package handler

import (
	handler "a6-api/handlers"
	model "a6-api/models"
	photo "a6-api/models/v1/photo"
	"a6-api/utils/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var subjectModel *photo.Subject

type mStruct struct {
	Subject string
	Nums    uint32
}

//SubjectList return subject list
func SubjectList(c *gin.Context) {
	baseParamsStruct := model.BaseParams{}
	listParamsStruct := photo.SubjectListParams{}
	baseParamsMaps := gin.H{}
	listParamsMaps := gin.H{}

	if c.ShouldBind(&baseParamsStruct) != nil || c.ShouldBind(&listParamsStruct) != nil {
		handler.ErrorMsg(c, http.StatusBadRequest, "")
	}

	baseParamsMaps = helper.Struct2Map(baseParamsStruct)
	helper.ParamTypeCovert(baseParamsMaps)
	listParamsMaps = helper.Struct2Map(listParamsStruct)

	data, pagin, notFound := subjectModel.List(baseParamsMaps, listParamsMaps)
	if notFound {
		handler.ErrorMsg(c, http.StatusNotFound, "not found record")
		return
	}
	handler.Ok(c, data, pagin, baseParamsStruct.ResponseType)
}

//SubjectDetail return subject detail
func SubjectDetail(c *gin.Context) {
	baseParamsStruct := model.BaseParams{}
	detailParamsStruct := photo.DetailParams{}
	detailParamsMaps := make(map[string]interface{})

	if c.ShouldBind(&baseParamsStruct) != nil {
		handler.ErrorMsg(c, http.StatusBadRequest, "")
	}

	detailParamsMaps = helper.Struct2Map(detailParamsStruct)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.ErrorMsg(c, http.StatusBadRequest, "Invalid params")
		return
	}

	detailParamsMaps["id"] = id
	data, err := subjectModel.Detail(detailParamsMaps)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			handler.ErrorMsg(c, http.StatusNotFound, err.Error())
			return
		}
		handler.ErrorMsg(c, http.StatusBadRequest, "Unknown error")
	} else {
		handler.Ok(c, data, map[string]interface{}{}, baseParamsStruct.ResponseType)
	}
}
