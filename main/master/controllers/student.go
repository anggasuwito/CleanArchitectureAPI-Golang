package controllers

import (
	"encoding/json"
	"gomux/main/master/models"
	"gomux/main/master/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

//StudentHandler StudentHandler
type StudentHandler struct {
	StudentUseCase usecases.StudentUseCase
}

//StudentController StudentController
func StudentController(r *mux.Router, service usecases.StudentUseCase) {
	StudentHandler := StudentHandler{service}
	r.HandleFunc("/students", StudentHandler.ListStudents).Methods(http.MethodGet)
	r.HandleFunc("/student/{id}", StudentHandler.StudentByID).Methods(http.MethodGet)
	r.HandleFunc("/student/{id}", StudentHandler.DeleteByID).Methods(http.MethodDelete)
	r.HandleFunc("/student/{id}", StudentHandler.UpdateStudent).Methods(http.MethodPut)
	r.HandleFunc("/student", StudentHandler.InsertStudent).Methods(http.MethodPost)
}

//ListStudents ListStudents
func (s StudentHandler) ListStudents(w http.ResponseWriter, r *http.Request) {
	students, err := s.StudentUseCase.GetStudents()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfStudents, err := json.Marshal(students)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfStudents)
}

//StudentByID StudentById
func (s StudentHandler) StudentByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idStudent := param["id"]
	student, err := s.StudentUseCase.GetStudentsByID(idStudent)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfStudents, err := json.Marshal(student)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfStudents)
}

//DeleteByID DeleteById
func (s StudentHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idStudent := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&s)
	err := s.StudentUseCase.DeleteStudentsByID(idStudent)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		w.Write([]byte("Data Deleted"))
	}
}

//UpdateStudent UpdateStudent
func (s StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idStudent := param["id"]
	var changeStudent models.Student
	_ = json.NewDecoder(r.Body).Decode(&changeStudent)
	err := s.StudentUseCase.UpdateStudentsByID(idStudent, changeStudent)
	if err != nil {
		w.Write([]byte("Id Not Found"))
	} else {
		w.Write([]byte("Data Updated"))
	}

}

//InsertStudent InsertStudent
func (s StudentHandler) InsertStudent(w http.ResponseWriter, r *http.Request) {
	var newStudent models.Student
	_ = json.NewDecoder(r.Body).Decode(&newStudent)
	err := s.StudentUseCase.InsertStudents(newStudent)
	if err != nil {
		w.Write([]byte("Insert Failed"))
	} else {
		w.Write([]byte("Insert Success"))
	}

}
