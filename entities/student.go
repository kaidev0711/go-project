package entities

import (
	"github.com/google/uuid"
	"github.com/kaidev0711/go-project/entities/shared"
)

type Student struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Age      int       `json:"age"`
}

var Students = []Student{
	{ID: shared.GetUuid(), FullName: "Tuan", Age: 22},
	{ID: shared.GetUuid(), FullName: "Thin", Age: 23},
}

func NewStudent(fullName string, age int) Student {
	return Student{
		ID:       shared.GetUuid(),
		FullName: fullName,
		Age:      age,
	}
}
