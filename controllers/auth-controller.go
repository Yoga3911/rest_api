package controllers

import (
	"rest_api/helper"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (a *authController) Login(c *fiber.Ctx) error {
	return c.JSON(helper.BuildResponse("success", true, "Login"))
}

func (a *authController) Register(c *fiber.Ctx) error {
	return c.JSON(helper.BuildResponse("success", true, "Register"))
}
