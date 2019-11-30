package middleware

import (
	"a6-api/utils/verification"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

func Validator() gin.HandlerFunc {
	return func(c *gin.Context) {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("subjectOrderFieldValid", verification.SubjectOrderFieldValid)
			v.RegisterValidation("orderTypeValid", verification.OrderTypeValid)
			v.RegisterValidation("responseTypeValid", verification.ResponseTypeValid)
		}
	}
}
