package handler

import (
	handler "a6-api/handlers"
	model "a6-api/models"
	photo "a6-api/models/v1/photo"
	"a6-api/utils/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

var housetypeModel *photo.Housetype

func HousetypeList(c *gin.Context) {
	baseParamsStruct := model.BaseParams{}
	baseParamsMaps := map[string]interface{}{}
	listParamsStruct := photo.HousetypeListParams{}
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

	data, pagin, notFound := housetypeModel.List(baseParamsMaps, listParamsMaps)
	if notFound {
		handler.ErrorMsg(c, http.StatusNotFound, "not found record", []string{})
		return
	}
	handler.Ok(c, data, pagin, baseParamsStruct.ResponseType)
}
