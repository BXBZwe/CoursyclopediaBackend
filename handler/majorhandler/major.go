package majorhandler

import (
	"BackendCoursyclopedia/service/majorservice"
)

type IMajorHandler interface {
}

type MajorHandler struct {
	MajorService majorservice.IMajorService
}

func NewMajorHandler(majorService majorservice.IMajorService) *MajorHandler {
	return &MajorHandler{
		MajorService: majorService,
	}
}
