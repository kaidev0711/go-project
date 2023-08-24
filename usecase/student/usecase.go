package student

import "github.com/kaidev0711/go-project/infra/database"

type StudentUsecase struct {
	Database *database.Database
}

func NewStudentUsecase(db *database.Database) *StudentUsecase {
	return &StudentUsecase{
		Database: db,
	}
}
