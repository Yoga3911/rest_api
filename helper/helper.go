package helper

import "github.com/gofiber/fiber/v2"

func BuildResponse(c *fiber.Ctx, m string, s bool, d interface{}) error {
	return c.JSON(fiber.Map{
		"message": m,
		"status":  s,
		"data":    d,
	})
}
