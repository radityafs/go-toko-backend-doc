package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	var err error

	var username = os.Getenv("DB_USERNAME")
	var password = os.Getenv("DB_PASSWORD")
	var host = os.Getenv("DB_HOST")
	var port = os.Getenv("DB_PORT")
	var database = os.Getenv("DB_DATABASE")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected")
	Migrate()
}
