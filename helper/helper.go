package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Failed string
	Tag string
	Value interface{}
}

func ErrorHandler(user interface{}) []*ErrorResponse {
	var errors []*ErrorResponse 
	validated := validator.New()
	err := validated.Struct(user)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			var er ErrorResponse
			er.Failed = e.StructNamespace()
			er.Tag = e.Tag()
			er.Value = e.Param()
			errors = append(errors, &er)
		}
	}
	return errors
}

func BuildResponse(m interface{}, s bool, d interface{}) interface{} {
	return fiber.Map{
		"message": m,
		"status":  s,
		"data":    d,
	}
}
