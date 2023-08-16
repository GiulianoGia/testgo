package db

import (
	"fmt"
	"gotest/types"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB_USERNAME string
var DB_PASSWORD string
var DB_NAME string
var DB_HOST string
var DB_PORT string

var DB *gorm.DB

func InitDB() {
	initEnvVariables()
	DB = ConnectDB()
	autoMigrateStructs()
}

func initEnvVariables() {
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
}

func autoMigrateStructs() {
	DB.AutoMigrate(&types.Grocery{})
	DB.AutoMigrate(&types.User{})
}

func ConnectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connection to database")
	}
	return db
}
