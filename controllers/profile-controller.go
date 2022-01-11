package controllers

import "github.com/gofiber/fiber/v2"

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
	return c.SendString("Update")
}

func (p *profileController) DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Delete")
}
