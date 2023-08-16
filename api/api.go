package api

import "github.com/gin-gonic/gin"

type Service struct {
	*gin.Engine
}

// New Service
func NewService() *Service {
	return &Service{
		gin.Default(),
	}
}

func (s *Service) Start() {
	s.GetRoutes()
	s.Run()
}
