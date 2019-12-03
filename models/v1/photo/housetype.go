package photo

import (
	"a6-api/utils/helper"
	"fmt"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

type Housetype struct{}

type HousetypeListFields struct {
	Id   uint8
	Name string
}

type HousetypeListParams struct {
	IsDefault  string `form:"-" json:"is_default,omitempty" map:"field:is_default:1"`
	IsShow     string `form:"-" json:"is_show,omitempty" map:"field:is_show;default:1"`
	OrderField string `form:"orderField" json:"orderField" map:"field:orderField;default:order_sort" binding:"omitempty,houseTypeFieldValid"`
}

func (*Housetype) TableName() string {
	prefix := photo.TablePrefix()
	table := "house_type"
	return fmt.Sprintf("%s_%s", prefix, table)
}

func (*HousetypeListParams) Error(err interface{}) string {
	if e, ok := err.(*strconv.NumError); ok {
		return fmt.Sprintf("This value %s is illegal", e.Num)
	} else if errors, has := err.(validator.ValidationErrors); has {
		//Validator.v9's errors
		for _, e := range errors {
			if e.StructNamespace() == "HousetypeListParams.OrderField" {
				switch e.ActualTag() {
				case "OrderField":
					return "Param \"OrderField\" only is [\"id\", \"order_sort\"]"
				}
			}
		}
	}

	//Other errors
	return "Invalid params"
}

func (h *Housetype) List(baseParamsMaps map[string]interface{}, listParamsMaps map[string]interface{}) (fields []HousetypeListFields, pagin map[string]interface{}, notFound bool) {
	var totalNum uint32

	db = db.Table(h.TableName())
	db.Where(listParamsMaps).Count(&totalNum)

	//get pagin data
	totalPage, offset := helper.Paginator(totalNum, baseParamsMaps["perPageNum"].(uint16), baseParamsMaps["page"].(uint16))

	if err := db.Where(listParamsMaps).Offset(offset).Order(baseParamsMaps["orderField"].(string) + " " + baseParamsMaps["orderType"].(string)).Limit(baseParamsMaps["perPageNum"].(uint16)).Scan(&fields).Error; err != nil {
		return []HousetypeListFields{}, pagin, true
	}

	//get pagin info
	pagin = helper.GeneratePaginInfo(totalNum, totalPage, baseParamsMaps["page"].(uint16), baseParamsMaps["perPageNum"].(uint16), offset)
	return
}
