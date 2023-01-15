package api_implementation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project17/entity"
	"project17/implementation"
)

type Serv struct {
	Inner *implementation.Basics
}

var student_request = new(entity.Student)

func (s *Serv) AddHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Add handler is called")
	json.NewDecoder(r.Body).Decode(&student_request)
	fmt.Println(student_request)
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(s.Inner.AddStudents(student_request))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Wrong method!"))
	}
}
func (s *Serv) GetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Get handler is called")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s.Inner.GetStudents())

}
func (s *Serv) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete handler is called")
	json.NewDecoder(r.Body).Decode(&student_request)
	fmt.Println(student_request)
	if r.Method == http.MethodDelete {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(s.Inner.DeleteStudents(student_request.Id))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Wrong method!"))
	}
}

func (s *Serv) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Update handler is called")
	json.NewDecoder(r.Body).Decode(&student_request)
	fmt.Println(student_request)
	if r.Method == http.MethodPut {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(s.Inner.UpdateStudents(student_request.Id, student_request))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Wrong method!"))
	}
}

// func (s *Serv) FindHandler(w http.ResponseWriter, r *http.Request) {

// 	json.NewDecoder(r.Body).Decode(&student_request)

//}
