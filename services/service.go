package services

import (
	"context"
	"tugas-cli/internal/config"
	"tugas-cli/internal/utils"
	"tugas-cli/internal/api"
	"time"
	"fmt"
)

type Service struct {
	ctx context.Context
	log *utils.Logger
	api *api.API
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

func (s *Service) taskWeather() {
	for {
		weather, err := s.api.GetWeather()
		if err != nil {
			s.log.Add(TAG, "WeatherAPI failed:" + err.Error())
		} else {
			temp := fmt.Sprintf("Current temperature in %s: %.2f", weather.Location.Name, weather.Current.TempC)
			s.log.Add(TAG, temp)
		}
		time.Sleep(time.Minute)
	}
}
