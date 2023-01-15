package implementation

import "project17/entity"

// type StudentInterface interface {
// 	AddStudents(list *entity.StudentList) *entity.StudentList
// 	GetStudents() *entity.StudentList
// 	DeleteStudents(list *entity.StudentList, id string) *entity.StudentList
// 	UpdateStudents(list *entity.StudentList, id string) *entity.StudentList
// }

func (basic_info *Basics) AddStudents(student_info *entity.Student) *entity.StudentList {
	basic_info.Entity.List = append(basic_info.Entity.List, *student_info)
	return basic_info.Entity
}

func (basic_info *Basics) GetStudents() *entity.StudentList {
	return basic_info.Entity
}

func (basic_info *Basics) UpdateStudents(id string, student_info *entity.Student) *entity.StudentList {
	if found, index := basic_info.FindStudent(id); found {
		basic_info.Entity.List[index] = *student_info
	}
	return basic_info.Entity
}
func (basic_info *Basics) DeleteStudents(id string) *entity.StudentList {
	if found, index := basic_info.FindStudent(id); found {
		copy(basic_info.Entity.List[index:], basic_info.Entity.List[index+1:])
		basic_info.Entity.List = basic_info.Entity.List[:len(basic_info.Entity.List)-1]
	}
	return basic_info.Entity
}
func (basic_info *Basics) FindStudent(id string) (bool, int) {
	result := false
	index_result := -1
	for index, student := range basic_info.Entity.List {
		if id == student.Id {
			result = true
			index_result = index
			break
		}
	}
	return result, index_result
}
