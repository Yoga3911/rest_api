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

	app := fiber.New(fiber.Config{
		Prefork:           true,
		StreamRequestBody: true,
	})

	ticker := time.NewTicker(25 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				app.Get("/", func(c *fiber.Ctx) error {
					return c.JSON(fiber.Map{
						"message": "OK!",
					})
				})
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	routes.Route(app)
	p := os.Getenv("PORT")
	p = fmt.Sprintf(":%v", p)

	app.Listen(p)
}
