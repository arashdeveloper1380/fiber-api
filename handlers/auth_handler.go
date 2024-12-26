package handlers

import (
	"crud-api-fiber/database"
	"crud-api-fiber/models"
	"crud-api-fiber/requests"
	utils "crud-api-fiber/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
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

	errFirst := database.DB.First(&user, "email = ?", loginReq.Email).Error
	if errFirst != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": errFirst.Error(),
		})
	}

	isValid := utils.CheckPassword(loginReq.Password, user.Password)

	if !isValid {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "password incorrect",
		})
	}

	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["phone"] = user.Phone
	claims["address"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errJWT := utils.GenerateJWT(&claims)
	if errJWT != nil {
		return errJWT
	}

	return ctx.Status(200).JSON(fiber.Map{
		"token": token,
	})
}
