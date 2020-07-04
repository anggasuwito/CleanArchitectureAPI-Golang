package usecases

import (
	"gomux/main/master/models"
	"gomux/main/master/repositories"
	"log"
)

//StudentUsecaseImpl StudentUsecaseImpl
type StudentUsecaseImpl struct {
	studentRepo repositories.StudentRepository
}

//GetStudents GetStudents
func (s StudentUsecaseImpl) GetStudents() ([]*models.Student, error) {
	student, err := s.studentRepo.GetAllStudent()
	if err != nil {
		return nil, err
	}
	return student, nil
}

//GetStudentsByID GetStudentsByID
func (s StudentUsecaseImpl) GetStudentsByID(id string) (models.Student, error) {
	student, err := s.studentRepo.GetAllStudentByID(id)
	if err != nil {
		log.Fatal(err)
	}
	return student, nil
}

//DeleteStudentsByID DeleteStudentsByID
func (s StudentUsecaseImpl) DeleteStudentsByID(id string) error {
	err := s.studentRepo.DeleteDataStudentByID(id)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

//UpdateStudentsByID UpdateStudentsByID
func (s StudentUsecaseImpl) UpdateStudentsByID(id string, changeStudent models.Student) error {
	err := s.studentRepo.UpdateStudentData(id, changeStudent)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

//InsertStudents InsertStudents
func (s StudentUsecaseImpl) InsertStudents(newStudent models.Student) error {
	err := s.studentRepo.InsertStudentData(newStudent)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

//InitStudentUseCase InitStudentUseCase
func InitStudentUseCase(studentRepo repositories.StudentRepository) StudentUseCase {
	return &StudentUsecaseImpl{studentRepo}
}
