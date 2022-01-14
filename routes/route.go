package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"rest_api/config"
	"rest_api/controllers"
	"rest_api/services"
)

var (
	DB       *pgxpool.Pool                 = config.DatabaseConnection()
	jwtS     services.JWTService           = services.NewJWTService()
	userS    services.UserService          = services.NewUserService(DB)
	userC    controllers.UserController    = controllers.NewUserController(userS)
	authS    services.AuthService          = services.NewAuthService(DB, jwtS)
	authC    controllers.AuthController    = controllers.NewAuthController(authS)
	profileS services.ProfileService       = services.NewProfileService(DB, jwtS)
	profileC controllers.ProfileController = controllers.NewProfileController(profileS)
)

func Route(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", userC.GetAllUser)
	api.Get("/user/:id", userC.GetUser)

	api.Post("/auth/register", authC.Register)
	api.Post("/auth/login", authC.Login)

	api.Put("/profile/user", profileC.UpdateUser)
	api.Put("/profile/users", profileC.UpdateUserByToken)
	api.Delete("/profile/user", profileC.DeleteUser)
	api.Delete("/profile/users", profileC.DeleteUserByToken)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "OK!",
		})
	})
}
