package handler

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

//defined enum type
var ResponseTypes = map[string]bool{
	"json":  true,
	"jsonp": true,
	"xml":   true,
}

//defined enum type
var OrderTypes = map[string]bool{
	"asc":  true,
	"desc": true,
}

//Record wrap record struct
type Record map[string]interface{}

//wrapRecord wrap record
func wrapRecord(data interface{}, pagin map[string]interface{}) (record Record) {
	t := reflect.TypeOf(data)
	switch t.Kind() {
	case reflect.Slice, reflect.Array:
		record = Record{"list": data, "paginInfo": pagin}
	case reflect.Struct:
		record = Record{"detail": data}
	default:
		record = Record{"content": data}
	}
	return
}

//Ok found record returns ok
func Ok(c *gin.Context, data interface{}, pagin map[string]interface{}, responseType string) {

	switch responseType {
	case "json":
		c.SecureJSON(http.StatusOK, gin.H{
			"msg":  http.StatusText(http.StatusOK),
			"data": wrapRecord(data, pagin),
		})
	case "jsonp":
		c.JSONP(http.StatusOK, fmt.Sprintf("callback=%s", data))
	case "xml":
		c.XML(http.StatusOK, gin.H{
			"msg":  http.StatusText(http.StatusOK),
			"data": wrapRecord(data, pagin),
		})
	case "yaml":
		c.YAML(http.StatusOK, gin.H{
			"msg":  http.StatusText(http.StatusOK),
			"data": wrapRecord(data, pagin),
		})
	default:
		c.SecureJSON(http.StatusOK, gin.H{
			"msg":  http.StatusText(http.StatusOK),
			"data": wrapRecord(data, pagin),
		})
	}
}

//ErrorMsg returns error message by http status code and param extMsg
func ErrorMsg(c *gin.Context, httpStatusCode int, extMsg string, data []string) {
	var msg string
	if extMsg != "" {
		msg = extMsg
	} else {
		msg = http.StatusText(httpStatusCode)
	}
	c.SecureJSON(httpStatusCode, gin.H{
		"msg":  fmt.Sprintf("%s", msg),
		"data": data,
	})
	c.Abort()
}
