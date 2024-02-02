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
}
