package student

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kaidev0711/go-project/entities"
	"github.com/kaidev0711/go-project/entities/shared"
)

func SearchByID(id uuid.UUID) (student entities.Student, err error) {
	for _, studentElement := range entities.Students {
		if studentElement.ID == id {
			student = studentElement
		}
	}
	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Don't find id")
	}
	return student, err
}
