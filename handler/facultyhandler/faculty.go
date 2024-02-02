package facultyhandler

import (
	"BackendCoursyclopedia/model/facultymodel"

	facultysvc "BackendCoursyclopedia/service/facultyservice"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type IFacultyHandler interface {
	GetFaculties(c *fiber.Ctx) error
	CreateFaculty(c *fiber.Ctx) error
	UpdateFaculty(c *fiber.Ctx) error
	DeleteFaculty(c *fiber.Ctx) error
}

type FacultyHandler struct {
	FacultyService facultysvc.IFacultyService
}

func NewFacultyHandler(facultyService facultysvc.IFacultyService) IFacultyHandler {
	return &FacultyHandler{
		FacultyService: facultyService,
	}
}

func (h FacultyHandler) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func (h FacultyHandler) GetFaculties(c *fiber.Ctx) error {
	ctx, cancel := h.withTimeout()
	defer cancel()

	faculties, err := h.FacultyService.GetAllFaculties(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Faculties retrieved successfully",
		"data":    faculties,
	})
}

func (h FacultyHandler) CreateFaculty(c *fiber.Ctx) error {
	ctx, cancel := h.withTimeout()
	defer cancel()

	var faculty facultymodel.Faculty
	if err := c.BodyParser(&faculty); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdFaculty, err := h.FacultyService.CreateFaculty(ctx, faculty)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Faculty created successfully",
		"data":    createdFaculty,
	})
}

func (h FacultyHandler) UpdateFaculty(c *fiber.Ctx) error {
	ctx, cancel := h.withTimeout()
	defer cancel()

	facultyID := c.Params("id")
	var faculty facultymodel.Faculty
	if err := c.BodyParser(&faculty); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedFaculty, err := h.FacultyService.UpdateFaculty(ctx, facultyID, faculty)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Faculty updated successfully",
		"data":    updatedFaculty,
	})
}

func (h FacultyHandler) DeleteFaculty(c *fiber.Ctx) error {
	ctx, cancel := h.withTimeout()
	defer cancel()

	facultyID := c.Params("id")
	err := h.FacultyService.DeleteFaculty(ctx, facultyID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Faculty deleted successfully",
	})
}
