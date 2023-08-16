package api

import (
	infra_controller "github.com/kaidev0711/go-project/api/controller/infra"
	students_controller "github.com/kaidev0711/go-project/api/controller/students"
)

func (s *Service) GetRoutes() {
	s.GET("/health", infra_controller.Health)

	groupStudents := s.Group("students")
	groupStudents.GET("/", students_controller.List)
	groupStudents.POST("/", students_controller.Create)
	groupStudents.PUT("/:id", students_controller.Update)
	groupStudents.DELETE("/:id", students_controller.Delete)
	groupStudents.GET("/:id", students_controller.Details)
}
