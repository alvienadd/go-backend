package main

import (
	"log"
	"os"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	port:=os.Getenv("PORT")
	app:=fiber.New

}