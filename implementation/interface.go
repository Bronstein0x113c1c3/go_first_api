package implementation

import (
	"project17/entity"
)

type StudentInterface interface {
	AddStudents(student_info *entity.Student) *entity.StudentList
	GetStudents() *entity.StudentList
	DeleteStudents(id string) *entity.StudentList
	UpdateStudents(id string, student_info *entity.Student) *entity.StudentList
	FindStudents(id string) (bool, int)
}
