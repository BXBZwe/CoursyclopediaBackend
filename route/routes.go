package route

import (
	"BackendCoursyclopedia/db"
	"BackendCoursyclopedia/handler/facultyhandler"
	"BackendCoursyclopedia/handler/userhandler"
	"BackendCoursyclopedia/repository/facultyrepository"
	userrepo "BackendCoursyclopedia/repository/userrepository"
	"BackendCoursyclopedia/service/facultyservice"
	usersvc "BackendCoursyclopedia/service/userservice"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	db.ConnectDB()

	userRepository := userrepo.NewUserRepository(db.DB)
	facultyReposiotry := facultyrepository.NewFacultyRepository(db.DB)

	userService := usersvc.NewUserService(userRepository)
	facultyService := facultyservice.NewFacultyService(facultyReposiotry)

	userHandler := userhandler.NewUserHandler(userService)
	facultyHandler := facultyhandler.NewFacultyHandler(facultyService)

	userGroup := app.Group("/api/users")
	userGroup.Get("/getallusers", userHandler.GetUsers)

	faculyGroup := app.Group("/api/faculties")
	faculyGroup.Get("getallfaculties", facultyHandler.GetFaculties)
	faculyGroup.Post("createfaculty", facultyHandler.CreateFaculty)
	faculyGroup.Put("updatefaculty/:id", facultyHandler.UpdateFaculty)
	faculyGroup.Delete("deletefaculty/:id", facultyHandler.DeleteFaculty)

}
