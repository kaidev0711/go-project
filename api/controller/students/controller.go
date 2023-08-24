package students

import student_usecase "github.com/kaidev0711/go-project/usecase/student"

type StudentController struct {
	StudentUsecase *student_usecase.StudentUsecase
}


func NewStudentController(su *student_usecase.StudentUsecase) *StudentController {
	return &StudentController{
		StudentUsecase: su,
	}
}
