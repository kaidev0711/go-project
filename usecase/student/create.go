package student

import (
	"github.com/kaidev0711/go-project/entities"
)

func (su *StudentUsecase) Create(fullName string, age int) (entities.Student, error) {
	student := entities.NewStudent(fullName, age)

	err := su.Database.StudentRepository.Create(&student)

	return student, err
}
