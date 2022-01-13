package helper

import (
	"github.com/gofiber/fiber/v2"
)

func BuildResponse(c *fiber.Ctx, fs int, m interface{}, s bool, d interface{}) error {
	return c.Status(fs).JSON(fiber.Map{
		"message": m,
		"status":  s,
		"data":    d,
	})
}
