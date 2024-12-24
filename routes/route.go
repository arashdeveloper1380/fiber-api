package routes

import (
	"crud-api-fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App) {

	userHandler := handlers.UserHandler{}

	api := route.Group("/api")

	api.Get("/", userHandler.All)
	api.Get("getById/:id", userHandler.GetById)
	api.Post("/create", userHandler.Create)
	api.Put("update/:id", userHandler.Update)
	api.Delete("delete/:id", userHandler.Delete)

}
