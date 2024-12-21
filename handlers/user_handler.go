package handlers

import (
	"crud-api-fiber/database"
	"crud-api-fiber/models"
	"crud-api-fiber/requests"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

func (uh *UserHandler) All(ctx *fiber.Ctx) error {

	var users []models.User

	err := database.DB.Find(&users).Error

	if err != nil {
		err := ctx.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
		if err != nil {
			return err
		}
	}

	return ctx.JSON(users)
}

func (uh *UserHandler) Create(ctx *fiber.Ctx) error {
	user := new(requests.UserCreateRequest)

	err := ctx.BodyParser(user)

	if err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": errValidate.Error(),
		})
	}

	newUser := models.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	err = database.DB.Create(&newUser).Error
	if err != nil {
		err := ctx.Status(500).JSON(fiber.Map{
			"error": "failed created user",
		})
		if err != nil {
			return err
		}
	}

	return ctx.JSON(fiber.Map{
		"message": "created successful",
		"data":    newUser,
	})

}

func (uh *UserHandler) GetById(ctx *fiber.Ctx) error {
	return nil
}
