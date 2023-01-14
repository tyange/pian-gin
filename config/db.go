package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"pian-gin/models"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlUserPassword := os.Getenv("MYSQL_PW")
	dbName := "test_burger4"

	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", mysqlUsername, mysqlUserPassword, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if db.AutoMigrate(&models.Burger{}, &models.User{}) != nil {
		log.Fatal("Failed DB auto migration.")
	}

	DB = db
}
