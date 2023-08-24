package student

import "github.com/kaidev0711/go-project/entities"

func (su *StudentUsecase) List() ([]entities.Student, error) {
	return su.Database.StudentRepository.List()
}
