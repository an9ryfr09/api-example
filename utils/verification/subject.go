package verification

import (
	"a6-api/utils/helper"
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

//SubjectOrderFieldValid handler subject's validate order field in [id|main_sort|sub_sort|personal_sort|special_sort|zmt_sort]
func SubjectOrderFieldValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if field.IsZero() {
		return true
	}

	orderFields := []string{
		"id", "main_sort", "sub_sort", "personal_sort", "special_sort", "zmt_sort",
	}
	return helper.IsContain(field.Interface().(string), orderFields, true)
}
