package main

import (
	"fmt"
	"os"
	"rest_api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	defer routes.DB.Close()
	
	app := fiber.New(fiber.Config{
		Prefork:           true,
		StreamRequestBody: true,
	})

	routes.Route(app)
	p := os.Getenv("PORT")
	p = fmt.Sprintf(":%v", p)

	app.Listen(p)
}
