package student

import "github.com/kaidev0711/go-project/entities"

func List() (students []entities.Student, err error) {
	students = entities.Students
	return students, err
}
