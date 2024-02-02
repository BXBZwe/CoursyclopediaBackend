package route

import (
	"BackendCoursyclopedia/db"
	"BackendCoursyclopedia/handler/auditloghandler"
	"BackendCoursyclopedia/handler/userhandler"

	auditlogrepo "BackendCoursyclopedia/repository/auditlogrepository"
	userrepo "BackendCoursyclopedia/repository/userrepository"

	auditlogsvc "BackendCoursyclopedia/service/auditlogservice"
	usersvc "BackendCoursyclopedia/service/userservice"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	db.ConnectDB()

	userRepo := userrepo.NewUserRepository(db.DB)
	auditlogRepo := auditlogrepo.NewAuditLogRepository(db.DB)

	userService := usersvc.NewUserService(userRepo)
	auditlogService := auditlogsvc.NewAuditLogService(auditlogRepo)

	userHandler := userhandler.NewUserHandler(userService)
	auditlogHandler := auditloghandler.NewAuditLogHandler(auditlogService)

	userGroup := app.Group("/api/users")
	userGroup.Get("/getallusers", userHandler.GetUsers)
	userGroup.Get("/getoneuser/:id", userHandler.GetOneUser)
	userGroup.Post("/createoneuser", userHandler.CreateOneUser)
	userGroup.Delete("/deleteoneuser/:id", userHandler.DeleteOneUser)
	userGroup.Put("/updateoneuser/:id", userHandler.UpdateOneUser)

	auditlogGroup := app.Group("/api/auditlogs")
	auditlogGroup.Get("/getallauditlogs", auditlogHandler.GetAuditLogs)

}
