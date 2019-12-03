package verification

import (
	"a6-api/utils/helper"

	"gopkg.in/go-playground/validator.v9"
)

//HouseTypeFieldValid handler housetype's validate order field in [id|order_sort]
func HouseTypeFieldValid(fl validator.FieldLevel) bool {
	if fl.Field().IsZero() {
		return true
	}

	orderFields := []string{
		"id", "order_sort",
	}
	return helper.IsContain(fl.Field().String(), orderFields, true)
}
