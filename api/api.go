package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kaidev0711/go-project/api/controller/students"
	"github.com/kaidev0711/go-project/infra/config"
	"github.com/kaidev0711/go-project/infra/database"
	student_usecase "github.com/kaidev0711/go-project/usecase/student"
)

type Service struct {
	Engine            *gin.Engine
	Database          *database.Database
	StudentController *students.StudentController
}

// New Service
func NewService(db *database.Database) *Service {
	return &Service{
		Engine:   gin.Default(),
		Database: db,
	}
}

func (s *Service) GetController() {
	studentUsecase := student_usecase.NewStudentUsecase(s.Database)
	s.StudentController = students.NewStudentController(studentUsecase)
}

func (s *Service) Start() error {
	s.GetController()
	s.GetRoutes()
	return s.Engine.Run(fmt.Sprintf(":%d", config.Env.ApiPort))
}
