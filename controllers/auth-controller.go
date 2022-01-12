package controllers

import (
	"rest_api/helper"
	"rest_api/models"
	"rest_api/services"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type authController struct {
	authS services.AuthService
}

func NewAuthController(authS services.AuthService) AuthController {
	return &authController{authS: authS}
}

func (a *authController) Login(c *fiber.Ctx) error {
	user := new(models.Login)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(helper.BuildResponse(err.Error(), false, nil))
	}

	errors := helper.ErrorHandler(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err2 := a.authS.VerifyCredential(c.Context(), *user)
	if err2 != nil {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(err2.Error(), false, nil))
	}
	return c.JSON(helper.BuildResponse("success", true, nil))
}

func (a *authController) Register(c *fiber.Ctx) error {
	user := new(models.Register)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(helper.BuildResponse(err.Error(), false, nil))
	}

	errors := helper.ErrorHandler(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err2 := a.authS.CreateUser(c.Context(), *user)
	if err2.Error() == "duplicate" {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(err2.Error(), false, nil))
	}

	return c.JSON(helper.BuildResponse("success", true, nil))
}
