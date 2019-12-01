package model

import (
	"fmt"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

type BaseParams struct {
	Page         uint16 `form:"page" json:"page" map:"field:page;default:1"`
	PerPageNum   uint16 `form:"perPageNum" json:"perPageNum" map:"field:perPageNum;default:10"`
	OrderType    string `form:"orderType" json:"orderType" map:"field:orderType;default:desc" binding:"orderTypeValid"`
	ResponseType string `form:"responseType" json:"responseType" map:"field:responseType;default:json" binding:"responseTypeValid"`
}

func (*BaseParams) Error(err interface{}) string {
	//Gin ShouldBindQuery's errors
	if e, ok := err.(*strconv.NumError); ok {
		return fmt.Sprintf("This value %s is illegal", e.Num)
	} else if errors, has := err.(validator.ValidationErrors); has {
		//Validator.v9's errors
		for _, e := range errors {
			if e.StructNamespace() == "BaseParams.OrderType" {
				switch e.ActualTag() {
				case "orderTypeValid":
					return "Param \"orderType\" only is [\"asc\" or \"desc\"]"
				}
			}
			if e.StructNamespace() == "BaseParams.ResponseType" {
				switch e.ActualTag() {
				case "responseTypeValid":
					return "Param \"responseType\" only is [\"json\", \"jsonp\", \"xml\", \"yaml\"]"
				}
			}
		}
	}
	//Other errors
	return "Invalid params"
}

func (*BaseParams) SwapParam(baseMap map[string]interface{}, listMap map[string]interface{}) {
	baseMap["orderField"] = listMap["orderField"]
	delete(listMap, "orderField")
}
