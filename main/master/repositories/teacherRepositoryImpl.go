package repositories

import (
	"database/sql"
	"gomux/main/master/models"
	"log"
)

//TeacherRepoImpl TeacherRepoImpl
type TeacherRepoImpl struct {
	db *sql.DB
}

//GetAllTeacher GetAllTeacher
func (s TeacherRepoImpl) GetAllTeacher() ([]*models.Teacher, error) {
	dataTeacher := []*models.Teacher{}
	query := "select * from teacher"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		teacher := models.Teacher{}
		var err = data.Scan(&teacher.TeacherID, &teacher.TeacherFname, &teacher.TeacherLname, &teacher.TeacherEmail)
		if err != nil {
			return nil, err
		}
		dataTeacher = append(dataTeacher, &teacher)
	}
	return dataTeacher, nil
}

//GetAllTeacherByID GetAllTeacherById
func (s TeacherRepoImpl) GetAllTeacherByID(id string) (models.Teacher, error) {
	var teacher models.Teacher
	query := "select * from teacher where teacher_id = ?"
	err := s.db.QueryRow(query, id).Scan(&teacher.TeacherID, &teacher.TeacherFname, &teacher.TeacherLname, &teacher.TeacherEmail)
	if err != nil {
		log.Fatal(err)
	}

	return teacher, nil
}

//DeleteDataTeacherByID DeleteDataTeacherById
func (s TeacherRepoImpl) DeleteDataTeacherByID(id string) error {
	tr, err := s.db.Begin()
	query := "delete from teacher where teacher_id = ?"
	_, err = s.db.Query(query, id)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}

	return nil
}

//UpdateTeacherData UpdateTeacherData
func (s TeacherRepoImpl) UpdateTeacherData(id string, changeTeacher models.Teacher) error {
	tr, err := s.db.Begin()
	query := "update teacher set teacher_fname=?,teacher_lname=?,teacher_email=? where teacher_id=?"
	_, err = s.db.Query(query, &changeTeacher.TeacherFname, &changeTeacher.TeacherLname, &changeTeacher.TeacherEmail, id)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}
	return nil
}

//InsertTeacherData InsertTeacherData
func (s TeacherRepoImpl) InsertTeacherData(newTeacher models.Teacher) error {

	tr, err := s.db.Begin()
	query := "insert into teacher (teacher_id,teacher_fname,teacher_lname,teacher_email) values (?,?,?,?)"
	_, err = s.db.Query(query, &newTeacher.TeacherID, &newTeacher.TeacherFname, &newTeacher.TeacherLname, &newTeacher.TeacherEmail)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}

	return nil
}

//InitTeacherRepoImpl InitTeacherRepoImpl
func InitTeacherRepoImpl(db *sql.DB) TeacherRepository {
	return &TeacherRepoImpl{db}
}
