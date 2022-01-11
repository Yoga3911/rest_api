package controllers

import (
	"rest_api/helper"
	"rest_api/models"
	"rest_api/services"

	"github.com/gofiber/fiber/v2"
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
	user := uc.userS.GetById(c.Context(), id)

	var u models.User
	err := user.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.GenderID, &u.CreateAt, &u.UpdateAt)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(err.Error(), false, nil))
	}
	return c.JSON(helper.BuildResponse("success", true, u))
}

func (uc *userController) GetAllUser(c *fiber.Ctx) error {
	users, err := uc.userS.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(helper.BuildResponse(err.Error(), false, nil))
	}
	var user []*models.User

	for users.Next() {
		var u models.User
		err2 := users.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.GenderID, &u.CreateAt, &u.UpdateAt)
		if err2 != nil {
			return c.Status(fiber.StatusConflict).JSON(helper.BuildResponse(err2.Error(), false, nil))
		}
		user = append(user, &u)
	}

	return c.JSON(helper.BuildResponse("success", true, user))
}
