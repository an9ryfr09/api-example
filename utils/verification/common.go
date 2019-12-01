package verification

import (
	"a6-api/utils/helper"

	"gopkg.in/go-playground/validator.v9"
)

//OrderTypeValid validate order type in [asc|desc]
func OrderTypeValid(fl validator.FieldLevel) bool {
	if fl.Field().IsZero() {
		return true
	}

	orderTypes := []string{
		"asc", "desc",
	}
	return helper.IsContain(fl.Field().String(), orderTypes, true)
}

//ResponseTypeValid validate response type in [json|jsonp|xml]
func ResponseTypeValid(fl validator.FieldLevel) bool {
	if fl.Field().IsZero() {
		return true
	}

	responseTypes := []string{
		"json", "jsonp", "xml", "yaml",
	}
	return helper.IsContain(fl.Field().String(), responseTypes, true)
}
