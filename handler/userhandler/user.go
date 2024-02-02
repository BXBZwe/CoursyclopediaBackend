package userhandler

import (
	usersvc "BackendCoursyclopedia/service/userservice"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type IUserHandler interface {
	GetUsers(c *fiber.Ctx) error
}

type UserHandler struct {
	UserService usersvc.IUserService
}

func NewUserHandler(userService usersvc.IUserService) IUserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h UserHandler) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	ctx, cancel := h.withTimeout()
	defer cancel()

	users, err := h.UserService.GetAllUsers(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Users retrieved successfully",
		"data":    users,
	})
}
