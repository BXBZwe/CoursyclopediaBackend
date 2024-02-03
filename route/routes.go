package route

import (
	"BackendCoursyclopedia/db"
	"BackendCoursyclopedia/handler/facultyhandler"
	"BackendCoursyclopedia/handler/majorhandler"
	"BackendCoursyclopedia/handler/userhandler"
	"BackendCoursyclopedia/repository/facultyrepository"
	"BackendCoursyclopedia/repository/majorrepository"
	userrepo "BackendCoursyclopedia/repository/userrepository"
	"BackendCoursyclopedia/service/facultyservice"
	"BackendCoursyclopedia/service/majorservice"
	usersvc "BackendCoursyclopedia/service/userservice"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	db.ConnectDB()

	userRepository := userrepo.NewUserRepository(db.DB)
	majorRepository := majorrepository.NewMajorRepository(db.DB)
	facultyRepository := facultyrepository.NewFacultyRepository(db.DB)

	userService := usersvc.NewUserService(userRepository)
	facultyService := facultyservice.NewFacultyService(facultyRepository)
	majorService := majorservice.NewMajorService(majorRepository, facultyRepository)

	userHandler := userhandler.NewUserHandler(userService)
	facultyHandler := facultyhandler.NewFacultyHandler(facultyService)
	majorHandler := majorhandler.NewMajorHandler(majorService)

	userGroup := app.Group("/api/users")
	userGroup.Get("/getallusers", userHandler.GetUsers)

	faculyGroup := app.Group("/api/faculties")
	faculyGroup.Get("getallfaculties", facultyHandler.GetFaculties)
	faculyGroup.Post("createfaculty", facultyHandler.CreateFaculty)
	faculyGroup.Put("updatefaculty/:id", facultyHandler.UpdateFaculty)
	faculyGroup.Delete("deletefaculty/:id", facultyHandler.DeleteFaculty)

	majorGroup := app.Group("api/majors")
	majorGroup.Post("createmajor", majorHandler.CreateMajor)
	majorGroup.Delete("deletemajor/:id", majorHandler.DeleteMajor)
	majorGroup.Put("updatemajor/:id", majorHandler.UpdateMajor)

}
