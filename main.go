package main

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var dbUser = os.Getenv("DB_USER")
var dbPassword = os.Getenv("DB_PASSWORD")
var dbName = os.Getenv("DB_NAME")
var dbHost = os.Getenv("DB_HOST")
var dbPort = os.Getenv("DB_PORT")
var usersFilePath = os.Getenv("USERS_PATH")

var db *gorm.DB
var dbErr error

func main() {

	fmt.Print("HELLO WORLD")

}
