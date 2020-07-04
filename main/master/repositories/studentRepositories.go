package repositories

import "gomux/main/master/models"

//StudentRepository StudentRepository
type StudentRepository interface {
	GetAllStudent() ([]*models.Student, error)
	GetAllStudentByID(id string) (models.Student, error)
	DeleteDataStudentByID(id string) error
	UpdateStudentData(id string, changeStudent models.Student) error
	InsertStudentData(models.Student) error
}
