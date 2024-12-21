package main

import (
	"crud-api-fiber/database"
	"crud-api-fiber/database/migrations"
	"crud-api-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	//Connect To DB
	database.DbInit()

	// Migrate Models To DB
	migrations.RunMigrate()

	// Boot Fiber App
	app := fiber.New()

	//Init Routes
	routes.RouteInit(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
