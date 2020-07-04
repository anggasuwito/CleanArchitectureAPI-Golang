package repositories

import "gomux/main/master/models"

//SubjectRepository SubjectRepository
type SubjectRepository interface {
	GetAllSubject() ([]*models.Subject, error)
	GetAllSubjectByID(id string) (models.Subject, error)
	DeleteDataSubjectByID(id string) error
	UpdateSubjectData(id string, changeSubject models.Subject) error
	InsertSubjectData(models.Subject) error
}
