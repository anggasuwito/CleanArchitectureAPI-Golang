package controllers

import (
	"encoding/json"
	"gomux/main/master/models"
	"gomux/main/master/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

//SubjectHandler SubjectHandler
type SubjectHandler struct {
	SubjectUseCase usecases.SubjectUseCase
}

//SubjectController SubjectController
func SubjectController(r *mux.Router, service usecases.SubjectUseCase) {
	SubjectHandler := SubjectHandler{service}
	r.HandleFunc("/subjects", SubjectHandler.ListSubjects).Methods(http.MethodGet)
	r.HandleFunc("/subject/{id}", SubjectHandler.SubjectByID).Methods(http.MethodGet)
	r.HandleFunc("/subject/{id}", SubjectHandler.DeleteByID).Methods(http.MethodDelete)
	r.HandleFunc("/subject/{id}", SubjectHandler.UpdateSubject).Methods(http.MethodPut)
	r.HandleFunc("/subject", SubjectHandler.InsertSubject).Methods(http.MethodPost)
}

//ListSubjects ListSubjects
func (s SubjectHandler) ListSubjects(w http.ResponseWriter, r *http.Request) {
	subjects, err := s.SubjectUseCase.GetSubjects()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfSubjects, err := json.Marshal(subjects)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfSubjects)
}

//SubjectByID SubjectById
func (s SubjectHandler) SubjectByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idSubject := param["id"]
	subject, err := s.SubjectUseCase.GetSubjectsByID(idSubject)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfSubjects, err := json.Marshal(subject)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfSubjects)
}

//DeleteByID DeleteById
func (s SubjectHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idSubject := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&s)
	err := s.SubjectUseCase.DeleteSubjectsByID(idSubject)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		w.Write([]byte("Data Deleted"))
	}
}

//UpdateSubject UpdateSubject
func (s SubjectHandler) UpdateSubject(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idSubject := param["id"]
	var changeSubject models.Subject
	_ = json.NewDecoder(r.Body).Decode(&changeSubject)
	err := s.SubjectUseCase.UpdateSubjectsByID(idSubject, changeSubject)
	if err != nil {
		w.Write([]byte("Id Not Found"))
	} else {
		w.Write([]byte("Data Updated"))
	}

}

//InsertSubject InsertSubject
func (s SubjectHandler) InsertSubject(w http.ResponseWriter, r *http.Request) {
	var newSubject models.Subject
	_ = json.NewDecoder(r.Body).Decode(&newSubject)
	err := s.SubjectUseCase.InsertSubjects(newSubject)
	if err != nil {
		w.Write([]byte("Insert Failed"))
	} else {
		w.Write([]byte("Insert Success"))
	}

}
