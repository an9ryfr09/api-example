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

func assertError(t *testing.T, got gin.H, want gin.H) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got type of %s, want type of %s", got, want)
	}
}

func TestStruct2Map(t *testing.T) {
	t.Run("struct to map:", func(t *testing.T) {

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

	t.Run("struct to map via json", func(t *testing.T) {
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
}
