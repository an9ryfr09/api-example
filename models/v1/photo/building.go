package photo

import (
	"a6-api/utils/helper"
	"fmt"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

type Building struct{}

type BuildingListFields struct {
	//主键
	Id uint64 `json:"id"`
	//楼盘名称
	Name string `json:"name"`
	//地址
	Address string `json:"address"`
	//省份id
	ProvinceId uint16 `json:"province"`
	//城市id
	CityId uint16 `json:"city"`
	//地区
	DistrictId uint32 `json:"district"`
	//经度
	Longitude float64 `json:"longitude"`
	//纬度
	Latitude float64 `json:"latitude"`
	//浏览量
	Views uint64 `json:"views"`
	//封面图
	Cover string `json:"cover"`
	//咨询数
	ConsultationNum uint64 `json:"consultation_num"`
	//签约数
	SignNum uint64 `json:"sign_num"`
	//开工户数
	StartNum uint64 `json:"start_num"`
	//竣工数
	CompletedNum uint64 `json:"completed_num"`
	//案例数
	CaseNum uint64 `json:"case_num"`
	//工地数
	ConstructionNum uint64 `json:"construction_num"`
	//楼盘类型(0=普通,1=推荐)
	Type uint8 `json:"type"`
	//标签(0=无, 1=特惠, 2=优惠, 4=热装)
	Label uint8 `json:"label"`
	//是否可参观样板间(0=否, 1=是)
	ModelroomVisitable uint8 `json:"modelroom_visitable"`
	//是否可参观工地(0=否, 1=是)
	ConstructionVisitable uint8 `json:"construction_visitable"`
	//是否可免费量房(0=否, 1=是)
	MeasuringHouse uint8 `json:"measuring_house"`
	//户型解析数
	HouseAnalysisNum uint64 `json:"house_analysis_num"`
	//服务站地址
	ServiceStationAddr string `json:"service_station_addr"`
	//附近实体店id
	NearbyStoreId uint64 `json:"nearby_store_id"`
}

type BuildingDetailFields struct {
	BuildingListFields
	//简介
	Description string `gorm:"description" json:"description"`
}

type BuildingListParams struct {
	SiteId     uint64 `form:"designerId" json:"designerId,omitempty" map:"field:designerid"`
	IsShow     string `form:"-" json:"is_show,omitempty" map:"field:is_show;default:yes"`
	OrderField string `form:"orderField" json:"orderField" binding:"buildingOrderFieldValid" map:"field:orderField;default:id"`
}

func (*BuildingListParams) Error(err interface{}) string {
	//Gin ShouldBindQuery's errors
	if e, ok := err.(*strconv.NumError); ok {
		return fmt.Sprintf("This value %s is illegal", e.Num)
	} else if errors, has := err.(validator.ValidationErrors); has {
		//Validator.v9's errors
		for _, e := range errors {
			if e.StructNamespace() == "BuildingListParams.OrderField" {
				switch e.ActualTag() {
				case "buildingOrderFieldValid":
					return "Param \"OrderField\" only is [\"id\", \"sort\"]"
				}
			}
		}
	}
	//Other errors
	return "Invalid params"
}

//TableName set this table name
func (*Building) TableName() string {
	prefix := photo.TablePrefix()
	table := "loupan"
	return fmt.Sprintf("%s_%s", prefix, table)
}

//List get query result for data list
func (b *Building) List(baseParamsMaps map[string]interface{}, listParamsMaps map[string]interface{}) (fields []BuildingListFields, pagin map[string]interface{}, notFound bool) {
	var totalNum uint32

	db = db.Table(b.TableName())
	db.Where(listParamsMaps).Count(&totalNum)

	//get pagin data
	totalPage, offset := helper.Paginator(totalNum, baseParamsMaps["perPageNum"].(uint16), baseParamsMaps["page"].(uint16))

	if err := db.Where(listParamsMaps).Offset(offset).Order(baseParamsMaps["orderField"].(string) + " " + baseParamsMaps["orderType"].(string)).Limit(baseParamsMaps["perPageNum"].(uint16)).Scan(&fields).Error; err != nil {
		return []BuildingListFields{}, pagin, true
	}

	//get pagin info
	pagin = helper.GeneratePaginInfo(totalNum, totalPage, baseParamsMaps["page"].(uint16), baseParamsMaps["perPageNum"].(uint16), offset)
	return
}

//Detail get query result for data detail
func (b *Building) Detail(detailParamsMaps map[string]interface{}) (fields BuildingDetailFields, err error) {
	if db.HasTable(b.TableName()) {
		err := db.Table(b.TableName()).Where(detailParamsMaps).Scan(&fields).Error
		return fields, err
	}
	return BuildingDetailFields{}, err
}
