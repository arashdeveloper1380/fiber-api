package migrations

import (
	"crud-api-fiber/database"
	"crud-api-fiber/models"
	"fmt"
	"log"
)

func RunMigrate() {
	if database.DB == nil {
		log.Fatalln("Database connection is not initialized")
	}

	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Migration successful")
}
