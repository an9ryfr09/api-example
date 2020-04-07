package photo

import (
	"a6-api/utils/helper"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type Designer struct{}

type DesignerListFields struct {
	//主键
	Id uint64 `gorm:"id" json:"id"`
	//设计师姓名
	Name string `gorm:"name" json:"name"`
	//姓名拼音
	Pinyin string `gorm:"pinyin" json:"pinyin"`
	//设计师头像
	Logo string `gorm:"logo" json:"logo"`
	//手机号
	Mobile string `gorm:"mobile" json:"mobile"`
	//报价
	Price float32 `gorm:"price" json:"price"`
	//qq
	Qq string `gorm:"qq" json:"qq"`
	//省份id
	ProvinceId uint16 `gorm:"province" json:"province"`
	//城市id
	CityId uint16 `gorm:"city" json:"city"`
	//分站id
	SiteId uint16 `gorm:"site_id" json:"site_id"`
	//门店id
	CompanyId uint16 `gorm:"company_id" json:"company_id"`
	//工作年限
	WorkYears uint8 `gorm:"work_years" json:"work_years"`
	//设计师级别
	Level uint8 `gorm:"level" json:"level"`
	//浏览量
	Views uint64 `gorm:"views" json:"views"`
	//设计师头衔(0=普通, 1=获奖设计师, 2=推荐设计师, 3=金牌设计师)
	Recommend uint8 `gorm:"recommend" json:"recommend"`
	//案例数
	CaseNum uint16 `gorm:"case_num" json:"case_num"`
	//关注数
	Favorites uint64 `gorm:"favorites" json:"favorites"`
	//户型解析数
	HouseAnalysisNum uint64 `gorm:"house_analysis_num" json:"house_analysis_num"`
	//设计师编号
	DesignerCode string `gorm:"designer_code" json:"designer_code"`
	//设计师类型(1=家装设计师, 2=软装设计师)
	Type uint8 `gorm:"type" json:"type"`
	//擅长风格
	Styles string `gorm:"styles" json:"styles"`
	//擅长空间类型
	Spaces string `gorm:"spaces" json:"spaces"`
	//擅长户型
	HouseTypes string `gorm:"house_types" json:"house_types"`
}

type DesignerDetailFields struct {
	DesignerListFields
	//获奖情况
	Awards string `gorm:"awards" json:"awards"`
	//自我介绍
	Description string `gorm:"description" json:"description"`
	//设计理念
	DesignIdea string `gorm:"design_idea" json:"design_idea"`
	//亮点
	Highlights string `gorm:"highlights" json:"highlights"`
}

type DesignerListParams struct {
	SiteId     uint64 `form:"designerId" json:"designerId,omitempty" map:"field:designerid"`
	IsShow     string `form:"-" json:"is_show,omitempty" map:"field:is_show;default:yes"`
	OrderField string `form:"orderField" json:"orderField" binding:"designerOrderFieldValid" map:"field:orderField;default:id"`
}

func (*DesignerListParams) Error(err interface{}) string {
	//Gin ShouldBindQuery's errors
	if e, ok := err.(*strconv.NumError); ok {
		return fmt.Sprintf("This value %s is illegal", e.Num)
	} else if errors, has := err.(validator.ValidationErrors); has {
		//Validator.v9's errors
		for _, e := range errors {
			if e.StructNamespace() == "DesignerListParams.OrderField" {
				switch e.ActualTag() {
				case "designerOrderFieldValid":
					return "Param \"OrderField\" only is [\"id\", \"sort\", \"sub_sort\", \"special_sort\"]"
				}
			}
		}
	}
	//Other errors
	return "Invalid params"
}

//TableName set this table name
func (*Designer) TableName() string {
	prefix := photo.TablePrefix()
	table := strings.ToLower(reflect.TypeOf(Designer{}).Name())
	return fmt.Sprintf("%s_%s", prefix, table)
}

//List ...
func (d *Designer) List(baseParamsMaps map[string]interface{}, listParamsMaps map[string]interface{}) (fields []DesignerListFields, pagin map[string]interface{}, notFound bool) {
	var totalNum uint32

	db = db.Table(d.TableName())
	db.Where(listParamsMaps).Count(&totalNum)

	//get pagin data
	totalPage, offset := helper.Paginator(totalNum, baseParamsMaps["perPageNum"].(uint16), baseParamsMaps["page"].(uint16))

	if err := db.Where(listParamsMaps).Offset(offset).Order(baseParamsMaps["orderField"].(string) + " " + baseParamsMaps["orderType"].(string)).Limit(baseParamsMaps["perPageNum"].(uint16)).Scan(&fields).Error; err != nil {
		return []DesignerListFields{}, pagin, true
	}

	//get pagin info
	pagin = helper.GeneratePaginInfo(totalNum, totalPage, baseParamsMaps["page"].(uint16), baseParamsMaps["perPageNum"].(uint16), offset)
	return
}

//Detail ...
func (d *Designer) Detail(detailParamsMaps map[string]interface{}) (fields DesignerDetailFields, err error) {
	if db.HasTable(d.TableName()) {
		err := db.Table(d.TableName()).Where(detailParamsMaps).Scan(&fields).Error
		return fields, err
	}
	return DesignerDetailFields{}, err
}
