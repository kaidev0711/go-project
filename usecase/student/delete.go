package student

import (
	"github.com/google/uuid"
	"github.com/kaidev0711/go-project/entities"
)

func Delete(id uuid.UUID) (err error) {
	var newStudent []entities.Student
	for _, studentElement := range entities.Students {
		if studentElement.ID == id {
			newStudent = append(newStudent, studentElement)  
		}
	}
	entities.Students = newStudent
	return err
	
}
