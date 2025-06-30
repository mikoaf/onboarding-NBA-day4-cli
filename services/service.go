package services

import (
	"context"
	"tugas-cli/internal/config"
	"tugas-cli/internal/utils"
	"time"
)

type Service struct {
	ctx context.Context
	log *utils.Logger
}

// Service instance creation function
func NewService(config *config.Config) *Service {
	logger := utils.NewLogger(10)
	return &Service{
		log: logger,
	}
}

// Just call this method to start your service instance and golang will do the rest
func (s *Service) Init(ctx context.Context) {
	s.ctx = ctx
	go s.taskOne()
	go s.taskTwo()
}

// Define all of your multitask application service here
func (s *Service) taskOne() {
	for {
		s.log.Add(TAG, "Service A")
		time.Sleep(500 * time.Millisecond)
	}
}

func (s *Service) taskTwo() {
	for {
		s.log.Add(TAG, "Service B")
		time.Sleep(3 * time.Second)
	}
}
