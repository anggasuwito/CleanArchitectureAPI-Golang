package controllers

import (
	"encoding/json"
	"gomux/main/master/models"
	"gomux/main/master/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

//TeacherHandler TeacherHandler
type TeacherHandler struct {
	TeacherUseCase usecases.TeacherUseCase
}

//TeacherController TeacherController
func TeacherController(r *mux.Router, service usecases.TeacherUseCase) {
	TeacherHandler := TeacherHandler{service}
	r.HandleFunc("/teachers", TeacherHandler.ListTeachers).Methods(http.MethodGet)
	r.HandleFunc("/teacher/{id}", TeacherHandler.TeacherByID).Methods(http.MethodGet)
	r.HandleFunc("/teacher/{id}", TeacherHandler.DeleteByID).Methods(http.MethodDelete)
	r.HandleFunc("/teacher/{id}", TeacherHandler.UpdateTeacher).Methods(http.MethodPut)
	r.HandleFunc("/teacher", TeacherHandler.InsertTeacher).Methods(http.MethodPost)
}

//ListTeachers ListTeachers
func (s TeacherHandler) ListTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := s.TeacherUseCase.GetTeachers()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfTeachers, err := json.Marshal(teachers)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfTeachers)
}

//TeacherByID TeacherById
func (s TeacherHandler) TeacherByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idTeacher := param["id"]
	teacher, err := s.TeacherUseCase.GetTeachersByID(idTeacher)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfTeachers, err := json.Marshal(teacher)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfTeachers)
}

//DeleteByID DeleteById
func (s TeacherHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idTeacher := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&s)
	err := s.TeacherUseCase.DeleteTeachersByID(idTeacher)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		w.Write([]byte("Data Deleted"))
	}
}

//UpdateTeacher UpdateTeacher
func (s TeacherHandler) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idTeacher := param["id"]
	var changeTeacher models.Teacher
	_ = json.NewDecoder(r.Body).Decode(&changeTeacher)
	err := s.TeacherUseCase.UpdateTeachersByID(idTeacher, changeTeacher)
	if err != nil {
		w.Write([]byte("Id Not Found"))
	} else {
		w.Write([]byte("Data Updated"))
	}

}

//InsertTeacher InsertTeacher
func (s TeacherHandler) InsertTeacher(w http.ResponseWriter, r *http.Request) {
	var newTeacher models.Teacher
	_ = json.NewDecoder(r.Body).Decode(&newTeacher)
	err := s.TeacherUseCase.InsertTeachers(newTeacher)
	if err != nil {
		w.Write([]byte("Insert Failed"))
	} else {
		w.Write([]byte("Insert Success"))
	}

}
