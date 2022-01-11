package controllers

import (
	"rest_api/helper"
	"rest_api/models"
	"rest_api/services"

	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type profileController struct {
	profileS services.ProfileService
}

func NewProfileController(profileS services.ProfileService) ProfileController {
	return &profileController{profileS: profileS}
}

func (p *profileController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(helper.BuildResponse(err.Error(), false, nil))
	}

	err2 := p.profileS.Update(c.Context(), *user, id)
	if err2 != nil {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(err2.Error(), false, nil))
	}
	return c.JSON(helper.BuildResponse("success", true, nil))
}

func (p *profileController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := p.profileS.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(err.Error(), false, nil))
	}
	return c.JSON(helper.BuildResponse("success", true, nil))
}
