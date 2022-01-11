package helper

import "github.com/gofiber/fiber/v2"

func BuildResponse(m interface{}, s bool, d interface{}) interface{} {
	return fiber.Map{
		"message": m,
		"status":  s,
		"data":    d,
	}
}
