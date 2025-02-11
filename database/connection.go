package database

import (
	// "JWT-Authentication-go/config"
	"JWT-Authentication-go/models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:Elias/096031@tcp(localhost:3306)/godb"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	db.AutoMigrate(& models.User{})

	return db, nil
}