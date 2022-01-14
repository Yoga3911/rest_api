package main

import (
	"github.com/gofiber/fiber/v2"
	"rest_api/routes"
)
//
func main() {
	defer routes.DB.Close()
	app := fiber.New(fiber.Config{
		Prefork:           true,
		StreamRequestBody: true,
	})
	routes.Route(app)
	app.Listen("127.0.0.1:8080")
}