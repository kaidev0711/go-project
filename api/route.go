package api

import (
	infra_controller "github.com/kaidev0711/go-project/api/controller/infra"
	students_controller "github.com/kaidev0711/go-project/api/controller/students"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/health", infra_controller.Health)

	groupStudents := s.Engine.Group("students")
	groupStudents.GET("/", s.StudentController.List)
	groupStudents.POST("/", s.StudentController.Create)
	groupStudents.PUT("/:id", students_controller.Update)
	groupStudents.DELETE("/:id", students_controller.Delete)
	groupStudents.GET("/:id", students_controller.Details)
}
