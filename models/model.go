package model

type BaseParams struct {
	Page         uint16 `form:"page" json:"page" map:"field:page;default:1"`
	PerPageNum   uint16 `form:"perPageNum" json:"perPageNum" map:"field:perPageNum;default:10"`
	OrderField   string `form:"orderField" json:"orderField" map:"field:orderField;default:id" binding:"subjectOrderFieldValid"`
	OrderType    string `form:"orderType" json:"orderType" map:"field:orderType;default:desc" binding:"orderTypeValid"`
	ResponseType string `form:"responseType" json:"responseType" map:"field:responseType;default:json" binding:"responseTypeValid"`
}
