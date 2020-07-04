package models

//Teacher Teacher
type Teacher struct {
	TeacherID    string `json:"teacher_id"`
	TeacherFname string `json:"teacher_fname"`
	TeacherLname string `json:"teacher_lname"`
	TeacherEmail string `json:"teacher_email"`
}
