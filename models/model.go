package model

//defined enum type
var ResponseTypes = map[string]bool{
	"json":  true,
	"jsonp": true,
}

var OrderTypes = map[string]bool{
	"asc":  true,
	"desc": true,
}

const (
	DEFAULT_ORDER_TYPE   = "desc"
	DEFAULT_ORDER_FIELDS = "sort"
)

type ListParams struct {
	Page         uint16 `form:"page"`
	PerPageNum   uint16 `form:"perPageNum"`
	OrderFields  string `form:"orderFields"`
	OrderType    string `form:"orderType"`
	ResponseType string `form:"responseType"`
}

type DetailParams struct {
	Id uint32 `json:"id" binding:"required" form:"id"`
}
