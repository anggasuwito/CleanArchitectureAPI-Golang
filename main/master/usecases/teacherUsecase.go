package usecases

import "gomux/main/master/models"

//TeacherUseCase TeacherUseCase
type TeacherUseCase interface {
	GetTeachers() ([]*models.Teacher, error)
	GetTeachersByID(id string) (models.Teacher, error)
	DeleteTeachersByID(id string) error
	UpdateTeachersByID(id string, changeTeacher models.Teacher) error
	InsertTeachers(models.Teacher) error
}
