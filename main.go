package main

import (
	"context"
	"log"

	"github.com/kaidev0711/go-project/api"
	"github.com/kaidev0711/go-project/infra/config"
	"github.com/kaidev0711/go-project/infra/database"
	"github.com/kaidev0711/go-project/infra/database/mongo"
	"github.com/kaidev0711/go-project/infra/database/mongo/repositories"
)

func main() {
	var err error
	ctx := context.TODO()
	err = config.StartConfig()
	FatalError(err)

	db := GetDatabase(ctx)

	err = api.NewService(db).Start()
	FatalError(err)
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetDatabase(ctx context.Context) *database.Database {
	client, err := mongo.GetConnection(ctx)
	FatalError(err)

	studentRepository := repositories.NewStudentRepository(ctx, client)

	return database.NewDatabase(client, studentRepository)
}
