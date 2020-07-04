package models

//Student Student
type Student struct {
	StudentID    string `json:"student_id"`
	StudentFname string `json:"student_fname"`
	StudentLname string `json:"student_lname"`
	StudentEmail string `json:"student_email"`
}
