package student

import (
	"github.com/kaidev0711/go-project/entities"
)

func Create(fullName string, age int) (student entities.Student, err error) {
	pointerStudent := entities.NewStudent(fullName, age)
	student = pointerStudent
	entities.Students = append(entities.Students, student)
	return student, err
}
