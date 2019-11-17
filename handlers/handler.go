package handler

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

//Record wrap record struct
type Record map[string]interface{}

//wrapRecord wrap record
func wrapRecord(data interface{}, pagin map[string]interface{}) (record Record) {
	var wrapContent string
	t := reflect.TypeOf(data)
	switch t.Kind() {
	case reflect.Slice, reflect.Array:
		wrapContent = "list"
		record = Record{wrapContent: data, "paginInfo": pagin}
	case reflect.Struct:
		wrapContent = "detail"
		record = Record{wrapContent: data}
	default:
		wrapContent = "content"
		record = Record{wrapContent: data}
	}
	return
}

//Ok found record returns ok
func Ok(c *gin.Context, data interface{}, pagin map[string]interface{}) {
	c.SecureJSON(http.StatusOK, gin.H{
		"msg":  http.StatusText(http.StatusOK),
		"data": wrapRecord(data, pagin),
	})
}

//ErrorMsg returns error message by http status code and param extMsg
func ErrorMsg(c *gin.Context, httpStatusCode int, extMsg string) {
	var msg string
	if extMsg != "" {
		msg = extMsg
	} else {
		msg = http.StatusText(httpStatusCode)
	}
	c.SecureJSON(httpStatusCode, gin.H{
		"msg":  fmt.Sprintf("%s", msg),
		"data": []string{},
	})
}
