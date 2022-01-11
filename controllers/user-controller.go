package controllers

import (
	"rest_api/helper"
	"rest_api/services"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUser(c *fiber.Ctx) error
	GetAllUser(c *fiber.Ctx) error
}

type userController struct {
	userS services.UserService
}

func NewUserController(userS services.UserService) UserController {
	return &userController{userS: userS}
}

func (uc *userController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := uc.userS.GetById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(err.Error(), false, nil))
	}

	return c.JSON(helper.BuildResponse("success", true, user))
}

func (uc *userController) GetAllUser(c *fiber.Ctx) error {
	users, err := uc.userS.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(err.Error(), false, nil))
	}

	return c.JSON(helper.BuildResponse("success", true, users))
}
