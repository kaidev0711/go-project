package repositories

import (
	"context"

	"github.com/kaidev0711/go-project/entities"
	mongo_infra "github.com/kaidev0711/go-project/infra/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongo_driver "go.mongodb.org/mongo-driver/mongo"
)

const (
	StudentCollection = "students"
)

type StudentRepository struct {
	Ctx        context.Context
	Collection *mongo_driver.Collection
}

func NewStudentRepository(ctx context.Context, client *mongo_driver.Client) *StudentRepository {
	return &StudentRepository{
		Ctx:        ctx,
		Collection: mongo_infra.GetCollection(ctx, client, StudentCollection),
	}
}

func (sr StudentRepository) Create(student *entities.Student) error {
	_, err := sr.Collection.InsertOne(sr.Ctx, student)
	return err
}

func (sr StudentRepository) List() (students []entities.Student, err error) {
	cur, err := sr.Collection.Find(sr.Ctx, bson.M{})
	if err != nil {
		return students, err
	}

	for cur.Next(sr.Ctx) {
		var student entities.Student

		if err = cur.Decode(&student); err != nil {
			return students, err
		}
		students = append(students, student)
	}
	return students, err
}
