package routes

import (
	"rest_api/controllers"

	"github.com/gofiber/fiber/v2"
)

var (
	userC controllers.UserController = controllers.NewUserController()
)

func Route(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", userC.GetAllUser)
	api.Get("/user/:id", userC.GetUser)

	// api.Post("/auth/register")
	// api.Post("/auth/login")

	// api.Put("/profile/user/:id")
	// api.Delete("/profile/user/:id")
}