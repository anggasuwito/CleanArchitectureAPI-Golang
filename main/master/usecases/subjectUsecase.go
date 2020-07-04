package usecases

import "gomux/main/master/models"

//SubjectUseCase SubjectUseCase
type SubjectUseCase interface {
	GetSubjects() ([]*models.Subject, error)
	GetSubjectsByID(id string) (models.Subject, error)
	DeleteSubjectsByID(id string) error
	UpdateSubjectsByID(id string, changeSubject models.Subject) error
	InsertSubjects(models.Subject) error
}
