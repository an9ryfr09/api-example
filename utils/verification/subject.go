package verification

import (
	"a6-api/utils/helper"

	"gopkg.in/go-playground/validator.v9"
)

//SubjectOrderFieldValid handler subject's validate order field in [id|main_sort|sub_sort|personal_sort|special_sort|zmt_sort]
func SubjectOrderFieldValid(fl validator.FieldLevel) bool {
	if fl.Field().IsZero() {
		return true
	}

	orderFields := []string{
		"id", "main_sort", "sub_sort", "personal_sort", "special_sort", "zmt_sort",
	}
	return helper.IsContain(fl.Field().String(), orderFields, true)
}
