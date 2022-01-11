package controllers

import (
	"rest_api/helper"

	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type profileController struct {
}

func NewProfileController() ProfileController {
	return &profileController{}
}

func (p *profileController) UpdateUser(c *fiber.Ctx) error {
	return c.JSON(helper.BuildResponse("success", true, "Update User"))
}

func (p *profileController) DeleteUser(c *fiber.Ctx) error {
	return c.JSON(helper.BuildResponse("success", true, "Delete User"))
}
