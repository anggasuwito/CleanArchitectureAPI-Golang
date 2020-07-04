package usecases

import (
	"gomux/main/master/models"
	"gomux/main/master/repositories"
	"log"
)

//TeacherUsecaseImpl TeacherUsecaseImpl
type TeacherUsecaseImpl struct {
	teacherRepo repositories.TeacherRepository
}

//GetTeachers GetTeachers
func (s TeacherUsecaseImpl) GetTeachers() ([]*models.Teacher, error) {
	teacher, err := s.teacherRepo.GetAllTeacher()
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

//GetTeachersByID GetTeachersByID
func (s TeacherUsecaseImpl) GetTeachersByID(id string) (models.Teacher, error) {
	teacher, err := s.teacherRepo.GetAllTeacherByID(id)
	if err != nil {
		log.Fatal(err)
	}
	return teacher, nil
}

//DeleteTeachersByID DeleteTeachersByID
func (s TeacherUsecaseImpl) DeleteTeachersByID(id string) error {
	err := s.teacherRepo.DeleteDataTeacherByID(id)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

//UpdateTeachersByID UpdateTeachersByID
func (s TeacherUsecaseImpl) UpdateTeachersByID(id string, changeTeacher models.Teacher) error {
	err := s.teacherRepo.UpdateTeacherData(id, changeTeacher)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

//InsertTeachers InsertTeachers
func (s TeacherUsecaseImpl) InsertTeachers(newTeacher models.Teacher) error {
	err := s.teacherRepo.InsertTeacherData(newTeacher)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

//InitTeacherUseCase InitTeacherUseCase
func InitTeacherUseCase(teacherRepo repositories.TeacherRepository) TeacherUseCase {
	return &TeacherUsecaseImpl{teacherRepo}
}
