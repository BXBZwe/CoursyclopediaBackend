package subjectservice

import (
	"BackendCoursyclopedia/model/subjectmodel"
	"BackendCoursyclopedia/repository/majorrepository"
	"BackendCoursyclopedia/repository/subjectrepository"
	"context"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"time"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type ISubjectService interface {
	CreateSubject(ctx context.Context, subject subjectmodel.Subject, majorId string) (string, error)
	DeleteSubject(subjectId string) error
	UpdateSubject(ctx context.Context, subjectId string, updates subjectmodel.SubjectUpdateRequest, newMajorId string) error
}

type SubjectService struct {
	SubjectRepository subjectrepository.ISubjectRepository
	MajorRepository   majorrepository.IMajorRepository
}

func NewSubjectService(SubjectRepo subjectrepository.ISubjectRepository, MajorRepo majorrepository.IMajorRepository) ISubjectService {
	return &SubjectService{
		SubjectRepository: SubjectRepo,
		MajorRepository:   MajorRepo,
	}
}

func (s *SubjectService) CreateSubject(ctx context.Context, subject subjectmodel.Subject, majorId string) (string, error) {
	subjectId, err := s.SubjectRepository.CreateSubject(ctx, subject)
	if err != nil {
		return "", err
	}

	if subject.SubjectStatus == "" {
		subject.SubjectStatus = "AVAILABLE" // Default status
	}

	// Convert subjectId (primitive.ObjectID) to a hexadecimal string
	subjectIdHex := subjectId.Hex()

	err = s.MajorRepository.AddSubjectToMajor(ctx, majorId, subjectIdHex)
	if err != nil {
		return "", err
	}

	return subjectIdHex, nil
}

func (s *SubjectService) DeleteSubject(subjectId string) error {
	ctx := context.Background()

	// Convert subjectId to ObjectID
	objId, err := primitive.ObjectIDFromHex(subjectId)
	if err != nil {
		return err
	}

	// Delete the subject from the subjects collection
	err = s.SubjectRepository.DeleteSubject(ctx, objId)
	if err != nil {
		return err
	}

	// Remove subject ID from all majors that contain it
	return s.MajorRepository.RemoveSubjectFromMajors(ctx, objId)
}

func (s *SubjectService) UpdateSubject(ctx context.Context, subjectId string, updates subjectmodel.SubjectUpdateRequest, newMajorId string) error {
	subjectObjId, err := primitive.ObjectIDFromHex(subjectId)
	if err != nil {
		return err // Error converting subject ID from hex
	}

	// Prepare the update document
	updateFields := bson.M{}

	if updates.SubjectCode != "" {
		updateFields["subjectCode"] = updates.SubjectCode
	}
	if updates.Name != "" {
		updateFields["name"] = updates.Name
	}
	if len(updates.Professors) > 0 {
		updateFields["professors"] = updates.Professors
	}
	if updates.SubjectDescription != "" {
		updateFields["subjectDescription"] = updates.SubjectDescription
	}
	if updates.Campus != "" {
		updateFields["campus"] = updates.Campus
	}
	if updates.Credit != nil {
		updateFields["credit"] = *updates.Credit
	}
	if updates.PreRequisite != nil {
		updateFields["pre_requisite"] = updates.PreRequisite
	}
	if updates.CoRequisite != nil {
		updateFields["co_requisite"] = updates.CoRequisite
	}
	if updates.SubjectStatus != "" {
		updateFields["subjectStatus"] = updates.SubjectStatus
	}
	if updates.AvailableDuration != nil {
		updateFields["available_duration"] = *updates.AvailableDuration
	}

	if newMajorId != "" {
		newmajObjId, err := primitive.ObjectIDFromHex(newMajorId)
		if err != nil {
			return err
		}

		currentmajor, err := s.MajorRepository.FindMajorBySubjectId(ctx, subjectObjId)
		if err != nil {
			return err
		}

		if currentmajor.ID != newmajObjId {
			err = s.MajorRepository.UpdatemajorforSubject(ctx, subjectObjId, currentmajor.ID, newmajObjId)
			if err != nil {
				return err
			}
		}
	}

	// // Perform the update if there are fields to update
	// if len(updateFields) > 0 {
	// 	err = s.SubjectRepository.UpdateSubject(ctx, subjectObjId, updateFields)
	// 	if err != nil {
	// 		return err // Error updating subject
	// 	}
	// }

	return nil
}
