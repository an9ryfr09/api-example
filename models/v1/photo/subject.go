package photo

import (
	model "a6-api/models"
	"a6-api/utils/helper"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Subject struct{}

//function list return fields
type SubjectListFields struct {
	//主键
	Id uint64 `json:"id"`
	//标题
	Subject string `json:"subject,omitempty" form:"subject"`
	//关键词
	Keywords string `json:"keywords"`
	//张数
	Nums uint32 `json:"nums"`
	//风格id
	StyleId uint8 `gorm:"column:style" json:"styleId"`
	//造价
	Cost uint32 `json:"cost"`
	//价格区间id
	PriceId uint8 `json:"priceId"`
	//户型id
	HouseTypeId uint8 `gorm:"column:housetype" json:"houseTypeId"`
	//户型图
	Housemap string `json:"housemap"`
	//设计师id
	DesignerId uint64 `gorm:"column:designerid" json:"designerId"`
	//省份id
	ProvinceId uint16 `gorm:"column:province" json:"provinceId"`
	//城市id
	CityId uint16 `gorm:"column:city" json:"cityId"`
	//封面图
	Cover string `gorm:"column:title_img" json:"cover"`
	//喜欢数
	LikeNum uint64 `gorm:"column:likes" json:"likeNum"`
	//ip地址
	IpAddress string `gorm:"column:ip" json:"ipAddress"`
	//添加日期
	AddDate time.Time `json:"addDate"`
	//添加用户
	AddUser string `json:"addUser"`
	//修改日期
	ModiDate time.Time `json:"modiDate"`
	//修改用户
	ModiUser string `json:"modiUser"`
	//楼盘id
	LoupanId uint16 `json:"loupanId"`
	//站点id
	SiteId uint16 `json:"siteId"`
	//店面id
	CompanyId uint16 `json:"companyId"`
	//工队id
	TeamId uint64 `json:"teamId"`
	//面积
	Area uint32 `json:"area"`
	//面积区间ID
	AreaId uint8 `json:"areaId"`
	//访问量
	VisitNum uint64 `gorm:"column:views" json:"visitNum"`
	//收藏量
	CollectCount uint64 `json:"collectCount"`
	//是否推荐(0=否,1=是)
	Recommend uint8 `json:"recommend"`
	//是否获奖(0=否,1=是)
	PrizeType uint8 `json:"prizeType"`
	//导入方式(0=自然,1=系统导入)
	ImportType uint8 `json:"importType"`
	//点击量
	HitNum uint64 `gorm:"column:hits" json:"hitNum"`
	//app点击量
	AppHits uint8 `json:"appHits"`
	//设计师状态(0=历届设计师,1=在职设计师)
	SpecialSort uint8 `json:"specialSort"`
	//自媒体案例排序
	ZmtSort uint8 `json:"zmtSort"`
	//案例类型(1=效果图,2=实景)
	Type uint8 `json:"type"`
	//装修类型(1=硬装,2=软装)
	DecorationType uint8 `json:"decorationType"`
	//监理id
	SupervisorId uint64 `json:"supervisorId"`
	//荣获奖项
	Award string `json:"award"`
}

type SubjectDetailParams struct {
	SubjectListFields
	//seo描述
	SEODescription string `gorm:"descriptions" json:"descriptions"`
	//户型图
	Housemap string `gorm:"housemap" json:"housemap"`
	//描述
	Description string `gorm:"description" json:"description"`
}

type SubjectListParams struct {
	DesignerId  uint64 `form:"designerId" json:"designerId,omitempty" map:"field:designerid"`
	HouseTypeId uint8  `form:"houseTypeId" json:"houseTypeId,omitempty" map:"field:housetype"`
	StyleId     uint8  `form:"styleId" json:"styleId,omitempty" map:"field:style"`
	AreaId      uint8  `form:"areaId" json:"areaId,omitempty" map:"field:area_id"`
	SiteId      uint8  `form:"siteId" json:"siteId,omitempty" map:"field:site_id"`
	Type        uint8  `form:"type" json:"type,omitempty" map:"field:type" binding:"omitempty,min=1,max=2"`
	IsShow      string `form:"-" json:"isshow,omitempty" map:"field:isshow;default:yes"`
	OrderField  string `form:"orderField" json:"orderField" map:"field:orderField;default:id" binding:"subjectOrderFieldValid"`
}

func (*SubjectListParams) Error(err interface{}) string {
	//Gin ShouldBindQuery's errors
	if e, ok := err.(*strconv.NumError); ok {
		return fmt.Sprintf("This value %s is illegal", e.Num)
	} else if errors, has := err.(validator.ValidationErrors); has {
		//Validator.v9's errors
		for _, e := range errors {
			if e.StructNamespace() == "SubjectListParams.OrderField" {
				switch e.ActualTag() {
				case "subjectOrderFieldValid":
					return "Param \"OrderField\" only is [\"id\", \"main_sort\", \"sub_sort\", \"personal_sort\", \"special_sort\", \"zmt_sort\"]"
				}
			}
			if e.StructNamespace() == "SubjectListParams.Type" {
				switch e.ActualTag() {
				case "min":
					return fmt.Sprintf("Param \"Type\" must greater than or equal to %s", e.Param())
				case "max":
					return fmt.Sprintf("Param \"Type\" must Less than or equal to %s", e.Param())
				}
			}
		}
	}
	//Other errors
	return "Invalid params"
}

type DetailParams struct {
	model.BaseParams
	Id     uint64 `map:"field:id"`
	IsShow string `map:"field:is_show;default:yes"`
}

var photo *Photo

//TableName set this table name
func (*Subject) TableName() string {
	prefix := photo.TablePrefix()
	table := strings.ToLower(reflect.TypeOf(Subject{}).Name())
	return fmt.Sprintf("%s_%s", prefix, table)
}

//List get query result for data list
func (s *Subject) List(baseParamsMaps map[string]interface{}, listParamsMaps map[string]interface{}) (fields []SubjectListFields, pagin map[string]interface{}, notFound bool) {
	var totalNum uint32

	db = db.Table(s.TableName())
	db.Where(listParamsMaps).Count(&totalNum)

	//get pagin data
	totalPage, offset := helper.Paginator(totalNum, baseParamsMaps["perPageNum"].(uint16), baseParamsMaps["page"].(uint16))

	if err := db.Where(listParamsMaps).Offset(offset).Order(baseParamsMaps["orderField"].(string) + " " + baseParamsMaps["orderType"].(string)).Limit(baseParamsMaps["perPageNum"].(uint16)).Scan(&fields).Error; err != nil {
		return []SubjectListFields{}, pagin, true
	}

	//get pagin info
	pagin = helper.GeneratePaginInfo(totalNum, totalPage, baseParamsMaps["page"].(uint16), baseParamsMaps["perPageNum"].(uint16), offset)
	return
}

//Detail get query result for data detail
func (s *Subject) Detail(detailParamsMaps map[string]interface{}) (fields SubjectDetailParams, err error) {
	if db.HasTable(s.TableName()) {
		err := db.Table(s.TableName()).Where(detailParamsMaps).Scan(&fields).Error
		return fields, err
	}
	return SubjectDetailParams{}, err
}
