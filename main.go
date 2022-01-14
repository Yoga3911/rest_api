package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"rest_api/routes"
	"time"
)

func main() {
	defer routes.DB.Close()

	for range time.Tick(time.Minute * 30) {
		fmt.Println("Ping!")
	}

	app := fiber.New(fiber.Config{
		Prefork:           true,
		StreamRequestBody: true,
	})

	routes.Route(app)
	p := os.Getenv("PORT")
	p = fmt.Sprintf(":%v", p)

	app.Listen(p)
}
