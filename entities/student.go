package entities

import (
	"github.com/google/uuid"
	"github.com/kaidev0711/go-project/entities/shared"
	// "github.com/kaidev0711/go-project/usecase/student"
)

type Student struct {
	ID       uuid.UUID `json:"id"        bson:"_id"`
	FullName string    `json:"full_name" bson:"full_name"`
	Age      int       `json:"age"       bson:"age"`
}

var Students = []Student{
	{ID: shared.GetUuid(), FullName: "Tuan", Age: 22},
	{ID: shared.GetUuid(), FullName: "Thin", Age: 23},
}

type StudentRepository interface {
	Create(student *Student) error
	List() ([]Student, error)
}

func NewStudent(fullName string, age int) Student {
	return Student{
		ID:       shared.GetUuid(),
		FullName: fullName,
		Age:      age,
	}
}
