package verification

import (
	"a6-api/utils/helper"
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

//OrderTypeValid validate order type in [asc|desc]
func OrderTypeValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if field.IsZero() {
		return true
	}

	orderTypes := []string{
		"asc", "desc",
	}
	return helper.IsContain(field.Interface().(string), orderTypes, true)
}

//ResponseTypeValid validate response type in [json|jsonp|xml]
func ResponseTypeValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if field.IsZero() {
		return true
	}

	responseTypes := []string{
		"json", "jsonp", "xml", "yaml",
	}
	return helper.IsContain(field.Interface().(string), responseTypes, true)
}
