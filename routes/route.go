package routes

import (
	"crud-api-fiber/handlers"
	"crud-api-fiber/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App) {

	route.Static("/public", "./public/asset")

	userHandler := handlers.UserHandler{}
	AuthHandler := handlers.AuthHandler{}

	api := route.Group("/api")

	api.Get("/", middlewares.XToken, userHandler.All)
	api.Get("getById/:id", userHandler.GetById)
	api.Post("/create", userHandler.Create)
	api.Put("update/:id", userHandler.Update)
	api.Delete("delete/:id", userHandler.Delete)

	api.Post("/login", middlewares.XToken, AuthHandler.Login)

}
