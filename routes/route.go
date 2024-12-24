package routes

import (
	"crud-api-fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func middleware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token != "secret" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	return ctx.Next()
}

func RouteInit(route *fiber.App) {

	route.Static("/public", "./public/asset")
	userHandler := handlers.UserHandler{}

	api := route.Group("/api")

	api.Get("/", middleware, userHandler.All)
	api.Get("getById/:id", userHandler.GetById)
	api.Post("/create", userHandler.Create)
	api.Put("update/:id", userHandler.Update)
	api.Delete("delete/:id", userHandler.Delete)

}
