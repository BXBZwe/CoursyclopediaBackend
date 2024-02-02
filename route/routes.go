package route

import (
	"BackendCoursyclopedia/db"
	"BackendCoursyclopedia/handler/userhandler"
	userrepo "BackendCoursyclopedia/repository/userrepository"
	usersvc "BackendCoursyclopedia/service/userservice"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	db.ConnectDB()

	userRepo := userrepo.NewUserRepository(db.DB)

	userService := usersvc.NewUserService(userRepo)

	userHandler := userhandler.NewUserHandler(userService)

	userGroup := app.Group("/api/users")
	userGroup.Get("/getallusers", userHandler.GetUsers)
	userGroup.Get("/getoneuser/:id", userHandler.GetOneUser)
	userGroup.Post("/createoneuser", userHandler.CreateOneUser)
	userGroup.Delete("/deleteoneuser/:id", userHandler.DeleteOneUser)
	userGroup.Put("/updateoneuser/:id", userHandler.UpdateOneUser)
}
