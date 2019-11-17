package model

const (
	ORDER_TYPE   = "desc"
	ORDER_FIELDS = "sort"
)

type Modeler interface {
	List() string
	Detail() string
}

type ListParams struct {
	Page        uint16 `form:"page"`
	PerPageNum  uint16 `form:"per_page_num"`
	OrderFields string `form:"order_fields"`
	OrderType   string `form:"order_type"`
}

type DetailParams struct {
	Id uint32 `json:"id" binding:"required" form:"id"`
}
