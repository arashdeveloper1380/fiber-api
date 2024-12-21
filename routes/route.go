package routes

import (
	"crud-api-fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App) {
	userHandler := handlers.UserHandler{}

	route.Get("/", userHandler.All)
	route.Get("/:id", userHandler.GetById)
	route.Post("/create", userHandler.Create)

	//route.Post("/create", func(ctx *fiber.Ctx) error {
	//	name := ctx.FormValue("name")
	//
	//	return ctx.JSON(fiber.Map{
	//		"name": name,
	//	})
	//})

}
