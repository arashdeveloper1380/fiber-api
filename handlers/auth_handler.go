package handlers

import (
	"crud-api-fiber/database"
	"crud-api-fiber/models"
	"crud-api-fiber/requests"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{}

func (au *AuthHandler) Login(ctx *fiber.Ctx) error {
	loginReq := new(requests.LoginRequest)

	err := ctx.BodyParser(loginReq)

	if err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginReq)
	if errValidate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": errValidate.Error(),
		})
	}

	var user models.User

	errFirst := database.DB.First(&user, "email = ? and password = ?", loginReq.Email, loginReq.Password).Error
	if errFirst != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": errFirst.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"token": user,
	})
}
