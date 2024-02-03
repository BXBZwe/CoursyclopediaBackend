package majorhandler

import (
	"BackendCoursyclopedia/service/majorservice"
	"context"

	"github.com/gofiber/fiber/v2"
)

type IMajorHandler interface {
	CreateMajor(c *fiber.Ctx) error
	DeleteMajor(c *fiber.Ctx) error
	UpdateMajor(c *fiber.Ctx) error
}

type MajorHandler struct {
	MajorService majorservice.IMajorService
}

func NewMajorHandler(majorService majorservice.IMajorService) *MajorHandler {
	return &MajorHandler{
		MajorService: majorService,
	}
}

func (h *MajorHandler) CreateMajor(c *fiber.Ctx) error {
	var request struct {
		MajorName string `json:"majorName"`
		FacultyID string `json:"facultyId"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.MajorService.CreateMajor(request.MajorName, request.FacultyID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Major created successfully",
	})
}

func (h *MajorHandler) DeleteMajor(c *fiber.Ctx) error {
	majorId := c.Params("id")

	err := h.MajorService.DeleteMajor(majorId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Major successfully deleted",
	})
}

func (h *MajorHandler) UpdateMajor(c *fiber.Ctx) error {
	majorId := c.Params("id")
	var request struct {
		NewMajorName string `json:"newMajorName"`
		NewFacultyID string `json:"newFacultyId"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ctx := context.Background()
	err := h.MajorService.UpdateMajor(ctx, majorId, request.NewMajorName, request.NewFacultyID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Major updated successfully",
	})
}
