package repositories

import (
	"database/sql"
	"gomux/main/master/models"
	"log"
)

//SubjectRepoImpl SubjectRepoImpl
type SubjectRepoImpl struct {
	db *sql.DB
}

//GetAllSubject GetAllSubject
func (s SubjectRepoImpl) GetAllSubject() ([]*models.Subject, error) {
	dataSubject := []*models.Subject{}
	query := "select * from subject"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		subject := models.Subject{}
		var err = data.Scan(&subject.SubjectID, &subject.SubjectName)
		if err != nil {
			return nil, err
		}
		dataSubject = append(dataSubject, &subject)
	}
	return dataSubject, nil
}

//GetAllSubjectByID GetAllSubjectById
func (s SubjectRepoImpl) GetAllSubjectByID(id string) (models.Subject, error) {
	var subject models.Subject
	query := "select * from subject where subject_id = ?"
	err := s.db.QueryRow(query, id).Scan(&subject.SubjectID, &subject.SubjectName)
	if err != nil {
		log.Fatal(err)
	}

	return subject, nil
}

//DeleteDataSubjectByID DeleteDataSubjectById
func (s SubjectRepoImpl) DeleteDataSubjectByID(id string) error {
	tr, err := s.db.Begin()
	query := "delete from subject where subject_id = ?"
	_, err = s.db.Query(query, id)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}

	return nil
}

//UpdateSubjectData UpdateSubjectData
func (s SubjectRepoImpl) UpdateSubjectData(id string, changeSubject models.Subject) error {
	tr, err := s.db.Begin()
	query := "update subject set subject_name=? where subject_id=?"
	_, err = s.db.Query(query, &changeSubject.SubjectName, id)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}
	return nil
}

//InsertSubjectData InsertSubjectData
func (s SubjectRepoImpl) InsertSubjectData(newSubject models.Subject) error {

	tr, err := s.db.Begin()
	query := "insert into subject (subject_id,subject_name) values (?,?)"
	_, err = s.db.Query(query, &newSubject.SubjectID, &newSubject.SubjectName)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}

	return nil
}

//InitSubjectRepoImpl InitSubjectRepoImpl
func InitSubjectRepoImpl(db *sql.DB) SubjectRepository {
	return &SubjectRepoImpl{db}
}
