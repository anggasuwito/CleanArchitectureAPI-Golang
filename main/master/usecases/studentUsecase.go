package usecases

import "gomux/main/master/models"

//StudentUseCase StudentUseCase
type StudentUseCase interface {
	GetStudents() ([]*models.Student, error)
	GetStudentsByID(id string) (models.Student, error)
	DeleteStudentsByID(id string) error
	UpdateStudentsByID(id string, changeStudent models.Student) error
	InsertStudents(models.Student) error
}
