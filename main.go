// main.go
package main

import (
	//"BackendCoursyclopedia/api"
	"BackendCoursyclopedia/db"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()
	defer db.DisconnectDB()

	app := fiber.New()

	//api.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
