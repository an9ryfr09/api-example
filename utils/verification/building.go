package verification

import (
	"a6-api/utils/helper"

	"gopkg.in/go-playground/validator.v9"
)

//BuildingOrderFieldValid handler subject's validate order field in [id|sort]
func BuildingOrderFieldValid(fl validator.FieldLevel) bool {
	if fl.Field().IsZero() {
		return true
	}

	orderFields := []string{
		"id", "sort",
	}
	return helper.IsContain(fl.Field().String(), orderFields, true)
}
