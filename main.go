// main.go
package main

import (
	"BackendCoursyclopedia/db"
	"log"

	"BackendCoursyclopedia/api"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// SetupRoutes configures the application routes
func SetupRoutes(app *fiber.App) {
	// Home Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to BackendCoursyclopedia API!")
	})
	// User routes
	app.Get("/api/users", api.GetUsers)
	app.Get("/api/users/:id", api.GetSpecificUser)
	app.Post("/api/users", api.NewUser)
	app.Put("/api/users/:id", api.UpdateUser)
	app.Delete("/api/users/:id", api.DeleteUser)

	// Add more routes for other resources as needed
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()
	defer db.DisconnectDB()

	app := fiber.New()
	SetupRoutes(app)

	port := "3000"

	// if port == ""{
	// 	port="3000"
	// }

	// log.Fatal(app.Listen(":3000"))
	log.Fatal(app.Listen("0.0.0.0:" + port))

}
