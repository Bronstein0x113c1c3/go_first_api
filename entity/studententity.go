package entity

type Student struct {
	Class Class  `json:"class"`
	Id    string `json:"id"`
	Name  string `json:"name"`
}
type StudentList struct {
	List []Student `json:"list"`
}
