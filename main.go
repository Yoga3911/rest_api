package main

import (
	"fmt"
	"os"
	"rest_api/routes"
	"time"

	"github.com/gofiber/fiber/v2"
)

//tes
func main() {
	defer routes.DB.Close()

	ticker := time.NewTicker(30 * time.Millisecond)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Hello master i'm still awake")
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
	p = fmt.Sprintf(":%v", p)
	app.Listen(p)
}
