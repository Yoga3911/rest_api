package controllers

import (
	"github.com/gofiber/fiber/v2"
	"rest_api/helper"
	"rest_api/models"
	"rest_api/services"
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
	var user models.Login

	err := c.BodyParser(&user)
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusNotAcceptable, err.Error(), false, nil)
	}

	errors := services.StructValidator(user)
	if errors != nil {
		return helper.BuildResponse(c, fiber.StatusBadRequest, errors, false, nil)
	}

	t, err := a.authS.VerifyCredential(c.Context(), user)
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusConflict, err.Error(), false, nil)
	}

	return helper.BuildResponse(c, fiber.StatusOK, "Login success", true, map[string]string{
		"token": t,
	})
}

func (a *authController) Register(c *fiber.Ctx) error {
	var user models.Register

	err := c.BodyParser(&user)
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusNotAcceptable, err.Error(), false, nil)
	}

	errors := services.StructValidator(user)
	if errors != nil {
		return helper.BuildResponse(c, fiber.StatusBadRequest, errors, false, nil)
	}

	err = a.authS.CreateUser(c.Context(), user)
	if err.Error() == "duplicate" {
		return helper.BuildResponse(c, fiber.StatusConflict, err.Error(), false, nil)
	}

	return helper.BuildResponse(c, fiber.StatusOK, "Register success", true, nil)
}
