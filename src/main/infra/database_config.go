package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"Fiber/src/main/app/api/models"
)

var DB *gorm.DB

func ConnectDatabase()  {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_go_fiber"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Product{})

	DB = database
}