package repositories

import "gomux/main/master/models"

//TeacherRepository TeacherRepository
type TeacherRepository interface {
	GetAllTeacher() ([]*models.Teacher, error)
	GetAllTeacherByID(id string) (models.Teacher, error)
	DeleteDataTeacherByID(id string) error
	UpdateTeacherData(id string, changeTeacher models.Teacher) error
	InsertTeacherData(models.Teacher) error
}
