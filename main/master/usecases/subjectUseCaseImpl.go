package usecases

import (
	"gomux/main/master/models"
	"gomux/main/master/repositories"
	"log"
)

//SubjectUsecaseImpl SubjectUsecaseImpl
type SubjectUsecaseImpl struct {
	subjectRepo repositories.SubjectRepository
}

//GetSubjects GetSubjects
func (s SubjectUsecaseImpl) GetSubjects() ([]*models.Subject, error) {
	subject, err := s.subjectRepo.GetAllSubject()
	if err != nil {
		return nil, err
	}
	return subject, nil
}

//GetSubjectsByID GetSubjectsByID
func (s SubjectUsecaseImpl) GetSubjectsByID(id string) (models.Subject, error) {
	subject, err := s.subjectRepo.GetAllSubjectByID(id)
	if err != nil {
		log.Fatal(err)
	}
	return subject, nil
}

//DeleteSubjectsByID DeleteSubjectsByID
func (s SubjectUsecaseImpl) DeleteSubjectsByID(id string) error {
	err := s.subjectRepo.DeleteDataSubjectByID(id)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

//UpdateSubjectsByID UpdateSubjectsByID
func (s SubjectUsecaseImpl) UpdateSubjectsByID(id string, changeSubject models.Subject) error {
	err := s.subjectRepo.UpdateSubjectData(id, changeSubject)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

//InsertSubjects InsertSubjects
func (s SubjectUsecaseImpl) InsertSubjects(newSubject models.Subject) error {
	err := s.subjectRepo.InsertSubjectData(newSubject)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

//InitSubjectUseCase InitSubjectUseCase
func InitSubjectUseCase(subjectRepo repositories.SubjectRepository) SubjectUseCase {
	return &SubjectUsecaseImpl{subjectRepo}
}
