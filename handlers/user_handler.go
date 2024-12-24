package handlers

import (
	"crud-api-fiber/database"
	"crud-api-fiber/models"
	"crud-api-fiber/requests"
	"crud-api-fiber/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
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

	hashPass, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(fiber.Map{
			"message": "not generate pass",
		})
	}
	newUser.Password = hashPass

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
	userId := ctx.Params("id")

	var user models.User

	err := database.DB.First(&user, userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"user": []string{},
		})
	}
	return ctx.JSON(fiber.Map{
		"data": user,
	})
}

func (uh *UserHandler) Update(ctx *fiber.Ctx) error {

	userId := ctx.Params("id")

	updateReq := new(requests.UpdateCreateRequest)

	err := ctx.BodyParser(updateReq)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var user models.User

	validate := validator.New()
	errValidate := validate.Struct(updateReq)

	if errValidate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": errValidate.Error(),
		})
	}

	errFind := database.DB.First(&user, userId).Error

	if errFind != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"data": []string{},
		})
	}

	if updateReq.Name != "" {
		user.Name = updateReq.Name
	}
	user.Address = updateReq.Address
	user.Phone = updateReq.Phone

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": errUpdate.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data": user,
	})
}

func (uh *UserHandler) Delete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user models.User

	errFound := database.DB.Debug().First(&user, userId).Error
	if errFound != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": errFound.Error(),
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": errDelete.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "user was deleted",
	})
}
