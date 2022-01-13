package controllers

import (
	"rest_api/helper"
	"rest_api/models"
	"rest_api/services"

	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	UpdateUser(c *fiber.Ctx) error
	UpdateUserByToken(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	DeleteUserByToken(c *fiber.Ctx) error
}

type profileController struct {
	profileS services.ProfileService
}

func NewProfileController(profileS services.ProfileService) ProfileController {
	return &profileController{profileS: profileS}
}

func (p *profileController) UpdateUser(c *fiber.Ctx) error {
	var user models.UpdateUser

	err := c.BodyParser(&user)
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusNotAcceptable, err.Error(), false, nil)
	}

	errors := services.StructValidator(user)
	if errors != nil {
		return helper.BuildResponse(c, fiber.StatusBadRequest, errors, false, nil)
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

	errors := services.StructValidator(user)
	if errors != nil {
		return helper.BuildResponse(c, fiber.StatusBadRequest, errors, false, nil)
	}

	err = p.profileS.Delete(c.Context(), user.ID)
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusConflict, err.Error(), false, nil)
	}

	return helper.BuildResponse(c, fiber.StatusOK, "Delete success", true, nil)
}

func (p *profileController) DeleteUserByToken(c *fiber.Ctx) error {
	err := p.profileS.DeleteByToken(c)
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusBadRequest, err.Error(), false, nil)
	}

	return helper.BuildResponse(c, fiber.StatusOK, "Delete success", true, nil)
}

func (p *profileController) UpdateUserByToken(c *fiber.Ctx) error {
	var user models.UpdateUser

	err := c.BodyParser(&user)
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusNotAcceptable, err.Error(), false, nil)
	}

	errors := services.StructValidator(user)
	if errors != nil {
		return helper.BuildResponse(c, fiber.StatusBadRequest, errors, false, nil)
	}

	err = p.profileS.UpdateByToken(c, user)
	if err != nil {
		return helper.BuildResponse(c, fiber.StatusConflict, err.Error(), false, nil)
	}

	return helper.BuildResponse(c, fiber.StatusOK, "Update success", true, nil)
}
