package repositories

import (
	"database/sql"
	"gomux/main/master/models"
	"log"
)

//StudentRepoImpl StudentRepoImpl
type StudentRepoImpl struct {
	db *sql.DB
}

//GetAllStudent GetAllStudent
func (s StudentRepoImpl) GetAllStudent() ([]*models.Student, error) {
	dataStudent := []*models.Student{}
	query := "select * from student"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		student := models.Student{}
		var err = data.Scan(&student.StudentID, &student.StudentFname, &student.StudentLname, &student.StudentEmail)
		if err != nil {
			return nil, err
		}
		dataStudent = append(dataStudent, &student)
	}
	return dataStudent, nil
}

//GetAllStudentByID GetAllStudentById
func (s StudentRepoImpl) GetAllStudentByID(id string) (models.Student, error) {
	var student models.Student
	query := "select * from student where student_id = ?"
	err := s.db.QueryRow(query, id).Scan(&student.StudentID, &student.StudentFname, &student.StudentLname, &student.StudentEmail)
	if err != nil {
		log.Fatal(err)
	}

	return student, nil
}

//DeleteDataStudentByID DeleteDataStudentById
func (s StudentRepoImpl) DeleteDataStudentByID(id string) error {
	tr, err := s.db.Begin()
	query := "delete from student where student_id = ?"
	_, err = s.db.Query(query, id)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}

	return nil
}

//UpdateStudentData UpdateStudentData
func (s StudentRepoImpl) UpdateStudentData(id string, changeStudent models.Student) error {
	tr, err := s.db.Begin()
	query := "update student set student_fname=?,student_lname=?,student_email=? where student_id=?"
	_, err = s.db.Query(query, &changeStudent.StudentFname, &changeStudent.StudentLname, &changeStudent.StudentEmail, id)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}
	return nil
}

//InsertStudentData InsertStudentData
func (s StudentRepoImpl) InsertStudentData(newStudent models.Student) error {

	tr, err := s.db.Begin()
	query := "insert into student (student_id,student_fname,student_lname,student_email) values (?,?,?,?)"
	_, err = s.db.Query(query, &newStudent.StudentID, &newStudent.StudentFname, &newStudent.StudentLname, &newStudent.StudentEmail)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}

	return nil
}

//InitStudentRepoImpl InitStudentRepoImpl
func InitStudentRepoImpl(db *sql.DB) StudentRepository {
	return &StudentRepoImpl{db}
}
