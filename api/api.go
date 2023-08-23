package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kaidev0711/go-project/infras/config"
)

type Service struct {
	*gin.Engine
}

// New Service
func NewService() *Service {
	return &Service{
		gin.Default(),
	}
}

func (s *Service) Start() error {
	s.GetRoutes()
	return s.Run(fmt.Sprintf(":%d", config.Env.ApiPort))
}
