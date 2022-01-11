package main

import (
	"rest_api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
		StreamRequestBody: true,
	})
	routes.Route(app)
	app.Listen("127.0.0.1:8080")
}