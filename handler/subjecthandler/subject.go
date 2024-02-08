package subjecthandler

import (
	"BackendCoursyclopedia/model/subjectmodel"
	"BackendCoursyclopedia/service/subjectservice"

	"github.com/gofiber/fiber/v2"
)

type ISubjectHandler interface {
	CreateSubject(c *fiber.Ctx) error
	DeleteSubject(c *fiber.Ctx) error
}
type SubjectHandler struct {
	SubjectService subjectservice.ISubjectService
}

func NewSubjectHandler(subjectService subjectservice.ISubjectService) *SubjectHandler {
	return &SubjectHandler{
		SubjectService: subjectService,
	}
}

func (h *SubjectHandler) CreateSubject(c *fiber.Ctx) error {
	var request struct {
		subjectmodel.Subject
		MajorId string `json:"majorId"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdSubjectId, err := h.SubjectService.CreateSubject(c.Context(), request.Subject, request.MajorId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": createdSubjectId, "message": "Major created successfully"})
}

func (h *SubjectHandler) DeleteSubject(c *fiber.Ctx) error {
	subjectId := c.Params("id")

	err := h.SubjectService.DeleteSubject(subjectId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Subject successfully deleted",
	})
}
