package main

import (
	"apotek-paten/Database"
	"apotek-paten/Model"
	"apotek-paten/Router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	Router.ServeApps()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {

	Database.Connect()
	Database.Database.AutoMigrate(&Model.User{})
}
