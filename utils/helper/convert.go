package helper

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//slice2MapWithContain slice to map via for enum type
//@param slice enum elements
//@param covertLower covert lower every elements
func slice2MapWithContain(slice []string, covertLower bool) map[string]bool {
	mapObj := map[string]bool{}
	value := true
	for _, v := range slice {
		if covertLower {
			v = strings.ToLower(v)
		}
		mapObj[v] = value
	}
	return mapObj
}

//Struct2Map is map to struct via reflect
//tips: embedded struct not supported
//@param structs struct elements
//usage:
//	type Foo struct{
//		Id uint32 `map:"field:id;default:1"`
//		Name string `map:field:name`
//		Age uint8 `map:field:age;default:25`
//	}
//
func Struct2Map(structs interface{}) (maps map[string]interface{}) {
	var tag, field string

	typeObj := reflect.TypeOf(structs)
	valueObj := reflect.ValueOf(structs)
	maps = gin.H{}

begin:
	for i := 0; i < typeObj.NumField(); i++ {

		tag = typeObj.Field(i).Tag.Get("map")

		if tag == "" {
			continue
		}

		if exists := strings.Contains(tag, ";"); !exists {
			elem := strings.Split(tag, ":")
			if elem[0] == "field" && !valueObj.Field(i).IsZero() {
				maps[elem[1]] = valueObj.Field(i).Interface()
			}
			continue
		}

		res := strings.Split(tag, ";")
		for j := 0; j < (len(res)); j++ {

			if exists := strings.Contains(tag, ":"); !exists {
				continue begin
			}

			elem := strings.Split(res[j], ":")

			switch elem[0] {
			case "field":
				field = elem[1]
				maps[field] = valueObj.Field(i).Interface()
			case "default":
				if valueObj.Field(i).IsZero() {
					maps[field] = elem[1]
				}
			case "empty":
				if b, _ := strconv.ParseBool(elem[1]); b {
					maps[field] = valueObj.Field(i).Interface()
				}
			}
		}
	}
	return
}

//Struct2MapViaJson is struct to map via json
//@structs struct elements
func Struct2MapViaJson(structs interface{}) (maps map[string]interface{}) {
	j, _ := json.Marshal(structs)
	json.Unmarshal(j, &maps)
	return
}

//ParamTypeCovert covert base params
//@param params base params sets
func ParamTypeCovert(params map[string]interface{}) {
	for k, p := range params {
		switch k {
		case "page", "perPageNum":
			switch p.(type) {
			case string:
				page, _ := strconv.ParseUint(p.(string), 10, 64)
				params[k] = uint16(page)
			case int:
				params[k] = uint16(p.(int))
			case uint16:
				params[k] = p
			}
		case "orderField", "orderType", "responseType":
			params[k] = p.(string)
		}
	}
}
