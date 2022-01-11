package routes

import (
	"rest_api/controllers"

	"github.com/gofiber/fiber/v2"
)

var (
	userC controllers.UserController = controllers.NewUserController()
	authC controllers.AuthController = controllers.NewAuthController()
)

func Route(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", userC.GetAllUser)
	api.Get("/user/:id", userC.GetUser)

	api.Post("/auth/register", authC.Register)
	api.Post("/auth/login", authC.Login)

	// api.Put("/profile/user/:id")
	// api.Delete("/profile/user/:id")
}