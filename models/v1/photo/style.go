package photo

import (
	"a6-api/utils/helper"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Style struct{}

type StyleListFields struct {
	Id      uint8
	Name    string
	Summary string
}

type StyleListParams struct {
	IsDefault  string `form:"-" json:"is_default,omitempty" map:"field:is_default:1"`
	IsShow     string `form:"-" json:"is_show,omitempty" map:"field:is_show;default:1"`
	OrderField string `form:"orderField" json:"orderField" map:"field:orderField;default:id"`
}

func (*Style) TableName() string {
	prefix := photo.TablePrefix()
	table := strings.ToLower(reflect.TypeOf(Style{}).Name())
	return fmt.Sprintf("%s_%s", prefix, table)
}

func (*StyleListParams) Error(err interface{}) string {
	if e, ok := err.(*strconv.NumError); ok {
		return fmt.Sprintf("This value %s is illegal", e.Num)
	}
	//Other errors
	return "Invalid params"
}

func (s *Style) List(baseParamsMaps map[string]interface{}, listParamsMaps map[string]interface{}) (fields []StyleListFields, pagin map[string]interface{}, notFound bool) {
	var totalNum uint32

	db = db.Table(s.TableName())
	db.Where(listParamsMaps).Count(&totalNum)

	//get pagin data
	totalPage, offset := helper.Paginator(totalNum, baseParamsMaps["perPageNum"].(uint16), baseParamsMaps["page"].(uint16))

	if err := db.Where(listParamsMaps).Offset(offset).Order(baseParamsMaps["orderField"].(string) + " " + baseParamsMaps["orderType"].(string)).Limit(baseParamsMaps["perPageNum"].(uint16)).Scan(&fields).Error; err != nil {
		return []StyleListFields{}, pagin, true
	}

	//get pagin info
	pagin = helper.GeneratePaginInfo(totalNum, totalPage, baseParamsMaps["page"].(uint16), baseParamsMaps["perPageNum"].(uint16), offset)
	return
}
