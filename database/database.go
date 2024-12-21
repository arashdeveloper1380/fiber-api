package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbInit() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/fiber-api?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can not connect ro database")
	}

	fmt.Println("Connected To The Database")
}
