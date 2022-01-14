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

	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fiber.Get("/")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	app := fiber.New(fiber.Config{
		Prefork:           true,
		StreamRequestBody: true,
	})
	
	routes.Route(app)
	p := os.Getenv("PORT")
	p = fmt.Sprintf(":%v",p)
	
	app.Listen(p)
}
