package controllers

import (
	"github.com/gofiber/fiber/v2"
	"rest_api/helper"
	"rest_api/services"
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
		return helper.BuildResponse(c, fiber.StatusConflict, err.Error(), false, nil)
	}

	return helper.BuildResponse(c, fiber.StatusOK, "Get user success", true, user)
}

func (uc *userController) GetAllUser(c *fiber.Ctx) error {
	users, err := uc.userS.GetAll(c.Context())
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusConflict, err.Error(), false, nil)
	}

	return helper.BuildResponse(c, fiber.StatusOK, "Get all user success", true, users)
}
