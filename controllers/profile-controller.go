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
	var user models.User
	
	err := c.BodyParser(&user)
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusNotAcceptable, err.Error(), false, nil)
	}
	
	err = p.profileS.Update(c.Context(), user)
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusConflict, err.Error(), false, nil)
	}
	
	return helper.BuildResponse(c, fiber.StatusOK, "Update success", true, nil)
}

func (p *profileController) DeleteUser(c *fiber.Ctx) error {	
	var user models.User
	
	err := c.BodyParser(&user)
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusNotAcceptable, err.Error(), false, nil)
	}
	
	err = p.profileS.Delete(c.Context(), user.ID)
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusConflict, err.Error(), false, nil)
	}

	return helper.BuildResponse(c, fiber.StatusOK, "Delete success", true, nil)
}
