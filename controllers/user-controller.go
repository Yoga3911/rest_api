package controllers

import (
	"rest_api/helper"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUser(c *fiber.Ctx) error
	GetAllUser(c *fiber.Ctx) error
}

type userController struct {

}

func NewUserController() UserController {
	return &userController{}
}

func (u *userController) GetUser(c *fiber.Ctx) error {
	return helper.BuildResponse(c, "success", true, "Get User")
}

func (u *userController) GetAllUser(c *fiber.Ctx) error {
	return helper.BuildResponse(c, "success", true, "Get All User")
}