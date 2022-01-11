package controllers

import (
	"rest_api/helper"
	"rest_api/models"
	"rest_api/services"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUser(c *fiber.Ctx) error
	GetAllUser(c *fiber.Ctx) error
}

type userController struct {
	userService services.UserService
}

func NewUserController(userS services.UserService) UserController {
	return &userController{userService: userS}
}

func (u *userController) GetUser(c *fiber.Ctx) error {
	return helper.BuildResponse(c, "success", true, "Get User")
}

func (u *userController) GetAllUser(c *fiber.Ctx) error {
	users, err := u.userService.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(c, err.Error(), false, nil))
	}
	var user []*models.User

	for users.Next() {
		var u models.User
		users.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.GenderID, &u.CreateAt, &u.UpdateAt)
		user = append(user, &u)
	}

	return helper.BuildResponse(c, "success", true, user)
}