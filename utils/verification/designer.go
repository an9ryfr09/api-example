package verification

import (
	"a6-api/utils/helper"

	"gopkg.in/go-playground/validator.v9"
)

//DesignerOrderFieldValid handler subject's validate order field in [id|sort]
func DesignerOrderFieldValid(fl validator.FieldLevel) bool {
	if fl.Field().IsZero() {
		return true
	}

	orderFields := []string{
		"id", "sort", "sub_sort", "special_sort",
	}
	return helper.IsContain(fl.Field().String(), orderFields, true)
}
