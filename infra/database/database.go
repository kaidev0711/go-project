package database

import (
	"github.com/kaidev0711/go-project/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Conn *mongo.Client

	StudentRepository entities.StudentRepository
}

func NewDatabase(conn *mongo.Client, sr entities.StudentRepository) *Database {
	return &Database{
		Conn: conn,
		StudentRepository: sr,
	}
}
