package model

import (
	"reflect"
	"testing"
)

func TestParamTypeCovert(t *testing.T) {
	t.Run("test param type covert:", func(t *testing.T) {
		var p *BaseParams
		maps := make(map[string]interface{})
		maps["page"] = "1"
		maps["perPageNum"] = "20"
		maps["responseType"] = "jsonp"
		p.ParamTypeCovert(maps)
		want := "uint16"
		got := reflect.TypeOf(maps["page"]).Kind()
		if want != got {
			t.Errorf("want: %s, got: %s", want, got)
		}
	})
}
