package config

import (
	"log"
	"myapp/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var JWT_SECRET = "1234"

func InitDB() {
	var err error
	dsn := "root:password@tcp(127.0.0.1:3306)/inventory_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established")

	// AutoMigrate for Users
	err = DB.AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
