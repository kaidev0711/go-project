package student

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kaidev0711/go-project/entities"
	"github.com/kaidev0711/go-project/entities/shared"
)

func Update(id uuid.UUID, fullName string, age int) (student entities.Student, err error) {
	var newStudent []entities.Student
	for _, studentElemt := range entities.Students {
		if studentElemt.ID == id {
			student = studentElemt
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Don't find student id")
	}

	student.FullName = fullName
	student.Age = age
	for _, studentElement := range entities.Students {
		if student.ID == studentElement.ID {
			newStudent = append(newStudent, student)
		} else {
			newStudent = append(newStudent, studentElement)
		}
	}
	entities.Students = newStudent

	return student, err
}
