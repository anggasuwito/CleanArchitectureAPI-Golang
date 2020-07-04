package master

import (
	"database/sql"
	"gomux/main/master/controllers"
	"gomux/main/master/repositories"
	"gomux/main/master/usecases"

	"github.com/gorilla/mux"
)

//Init Init
func Init(r *mux.Router, db *sql.DB) {
	studentRepo := repositories.InitStudentRepoImpl(db)
	studentUseCase := usecases.InitStudentUseCase(studentRepo)
	controllers.StudentController(r, studentUseCase)
	teacherRepo := repositories.InitTeacherRepoImpl(db)
	teacherUseCase := usecases.InitTeacherUseCase(teacherRepo)
	controllers.TeacherController(r, teacherUseCase)
	subjectRepo := repositories.InitSubjectRepoImpl(db)
	subjectUseCase := usecases.InitSubjectUseCase(subjectRepo)
	controllers.SubjectController(r, subjectUseCase)
}
