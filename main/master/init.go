package master

import (
	"gomux/config"
	"gomux/main/master/controllers"
	"gomux/main/master/repositories"
	"gomux/main/master/usecases"
)

//Init Init
func Init() {
	db := config.InitDB()
	router := config.CreateRouter()

	studentRepo := repositories.InitStudentRepoImpl(db)
	studentUseCase := usecases.InitStudentUseCase(studentRepo)
	controllers.StudentController(router, studentUseCase)
	teacherRepo := repositories.InitTeacherRepoImpl(db)
	teacherUseCase := usecases.InitTeacherUseCase(teacherRepo)
	controllers.TeacherController(router, teacherUseCase)
	subjectRepo := repositories.InitSubjectRepoImpl(db)
	subjectUseCase := usecases.InitSubjectUseCase(subjectRepo)
	controllers.SubjectController(router, subjectUseCase)

	config.RunServer(router)
}
