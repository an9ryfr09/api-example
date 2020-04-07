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

var designerModel *photo.Designer

func DesignerList(c *gin.Context) {
	baseParamsStruct := model.BaseParams{}
	baseParamsMaps := map[string]interface{}{}
	listParamsStruct := photo.DesignerListParams{}
	listParamsMaps := map[string]interface{}{}

	if err := c.ShouldBindQuery(&baseParamsStruct); err != nil {
		handler.ErrorMsg(c, http.StatusBadRequest, baseParamsStruct.Error(err), []string{})
		return
	}

	baseParamsMaps = helper.Struct2Map(baseParamsStruct)
	helper.ParamTypeCovert(baseParamsMaps)

	if err := c.ShouldBindQuery(&listParamsStruct); err != nil {
		handler.ErrorMsg(c, http.StatusBadRequest, listParamsStruct.Error(err), []string{})
		return
	}
	listParamsMaps = helper.Struct2Map(listParamsStruct)

	baseParamsStruct.SwapParam(baseParamsMaps, listParamsMaps)

	data, pagin, notFound := designerModel.List(baseParamsMaps, listParamsMaps)
	if notFound {
		handler.ErrorMsg(c, http.StatusNotFound, "not found record", []string{})
		return
	}
	handler.Ok(c, data, pagin, baseParamsStruct.ResponseType)
}

func DesignerDetail(c *gin.Context) {
	baseParamsStruct := model.BaseParams{}
	detailParamsStruct := photo.DetailParams{}
	detailParamsMaps := make(map[string]interface{})

	if c.ShouldBindQuery(&baseParamsStruct) != nil {
		handler.ErrorMsg(c, http.StatusBadRequest, "", []string{})
	}

	detailParamsMaps = helper.Struct2Map(detailParamsStruct)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.ErrorMsg(c, http.StatusBadRequest, "Invalid params", []string{})
		return
	}

	detailParamsMaps["id"] = id
	data, err := designerModel.Detail(detailParamsMaps)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			handler.ErrorMsg(c, http.StatusNotFound, err.Error(), []string{})
			return
		}
		handler.ErrorMsg(c, http.StatusBadRequest, "Unknown error", []string{})
	} else {
		handler.Ok(c, data, map[string]interface{}{}, baseParamsStruct.ResponseType)
	}
	return
}
