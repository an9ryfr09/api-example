package helper

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

type testStruct struct {
	Id   uint32 `map:"field:id;empty:true" json:"id,number"`
	Name string `map:"field:name;default:tom" json:"name,string"`
	Age  uint16 `map:"field:age;default:30" json:"age,number"`
}

func assertError(t *testing.T, got map[string]interface{}, want map[string]interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got type of %s, want type of %s", got, want)
	}
}

func TestCovert(t *testing.T) {
	t.Run("function Struct2Map", func(t *testing.T) {

		test := testStruct{
			Id:   0,
			Name: "yanlei",
			Age:  35,
		}

		got := Struct2Map(test)

		want := gin.H{
			"id":   uint32(0),
			"name": "yanlei",
			"age":  uint16(35),
		}

		assertError(t, got, want)
	})

	t.Run("function Struct2MapViaJson", func(t *testing.T) {
		test := testStruct{
			Id:   5,
			Name: "yanlei",
			Age:  35,
		}

		got := Struct2MapViaJson(test)

		want := gin.H{
			"id":   float64(5),
			"name": "\"yanlei\"",
			"age":  float64(35),
		}

		assertError(t, got, want)
	})

	t.Run("function ParamTypeCovert", func(t *testing.T) {
		maps := map[string]interface{}{}
		wantType := map[string]interface{}{}
		gotType := map[string]interface{}{}

		maps["page"] = "1"
		maps["perPageNum"] = "20"
		maps["responseType"] = "jsonp"

		ParamTypeCovert(maps)

		wantType["page"] = reflect.Uint16
		wantType["perPageNum"] = reflect.Uint16
		wantType["responseType"] = reflect.String

		gotType["page"] = reflect.TypeOf(maps["page"]).Kind()
		gotType["perPageNum"] = reflect.TypeOf(maps["perPageNum"]).Kind()
		gotType["responseType"] = reflect.TypeOf(maps["responseType"]).Kind()

		assertError(t, gotType, wantType)
	})
}
