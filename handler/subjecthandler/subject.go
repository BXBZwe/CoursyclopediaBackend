package subjecthandler

import (
	"BackendCoursyclopedia/model/subjectmodel"
	"BackendCoursyclopedia/service/subjectservice"
	"context"

	// "context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type ISubjectHandler interface {
	CreateSubject(c *fiber.Ctx) error
	DeleteSubject(c *fiber.Ctx) error
	UpdateSubject(c *fiber.Ctx) error
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

func (h *SubjectHandler) UpdateSubject(c *fiber.Ctx) error {
	subjectId := c.Params("id")
	var request struct {
		subjectmodel.SubjectUpdateRequest
		Professors []string `json:"professors"`
		NewMajorId string   `json:"newMajorId"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	professorObjectIDs := make([]primitive.ObjectID, len(request.Professors))
	for i, profStr := range request.Professors {
		profID, err := primitive.ObjectIDFromHex(profStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid professor ID"})
		}
		professorObjectIDs[i] = profID
	}

	request.SubjectUpdateRequest.Professors = professorObjectIDs

	ctx := context.Background()
	err := h.SubjectService.UpdateSubject(ctx, subjectId, request.SubjectUpdateRequest, request.NewMajorId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Subject updated successfully",
	})
}
