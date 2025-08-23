package main

import "crud-api-fiber/pkg"

//func main() {
//
//	//Connect To DB
//	database.DbInit()
//
//	// Migrate Models To DB
//	migrations.RunMigrate()
//
//	// Boot Fiber App
//	app := fiber.New()
//
//	//Init Routes
//	routes.RouteInit(app)
//
//	err := app.Listen(":3000")
//	if err != nil {
//		return
//	}
//}

func main() {
	pkg.CheckVariableWithReflect()

	pkg.CheckStructFieldReflect()

	pkg.CallFuncDynamicReflect()

	pkg.GetStructPropertyTags()

}
