package verification

import (
	"a6-api/utils/helper"

	"gopkg.in/go-playground/validator.v9"
)

//HousePriceFieldValid handler houseprice's validate order field in [id]
func HousePriceFieldValid(fl validator.FieldLevel) bool {
	if fl.Field().IsZero() {
		return true
	}

	orderFields := []string{
		"id",
	}
	return helper.IsContain(fl.Field().String(), orderFields, true)
}
