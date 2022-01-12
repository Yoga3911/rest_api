package routes

import (
	"rest_api/config"
	"rest_api/controllers"
	"rest_api/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	DB *pgxpool.Pool = config.DatabaseConnection()
	userS services.UserService = services.NewUserService(DB)
	userC controllers.UserController = controllers.NewUserController(userS)
	authS services.AuthService = services.NewAuthService(DB)
	authC controllers.AuthController = controllers.NewAuthController(authS)
	profileS services.ProfileService = services.NewProfileService(DB)
	profileC controllers.ProfileController = controllers.NewProfileController(profileS)
)

func Route(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", userC.GetAllUser)
	api.Get("/user/:id", userC.GetUser)

	api.Post("/auth/register", authC.Register)
	api.Post("/auth/login", authC.Login)

	api.Put("/profile/user", profileC.UpdateUser)
	api.Delete("/profile/user", profileC.DeleteUser)
}