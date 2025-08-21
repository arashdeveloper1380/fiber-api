package pkg

import (
	"crud-api-fiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func GetInstance() *gorm.DB {
	once.Do(func() {
		dsn := "root:@tcp(127.0.0.1:3306)/fiber-api?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // متغیر محلی db
		if err != nil {
			panic("Cannot connect to database")
		}

		instance = db
	})

	return instance
}

func userInstanceExample() {
	gormDB := GetInstance()
	var user models.User
	gormDB.Model(models.User{}).Where("id", 1).Scan(user)
}
