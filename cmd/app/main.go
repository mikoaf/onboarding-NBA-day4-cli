package main

import (
	"tugas-cli/internal/config"
	"tugas-cli/internal/utils"
	"tugas-cli/services"
	"fmt"
)

const TAG = "MAIN"
const MESSAGE = "Starting Application v1.0.0"
const CHANGELOG = "Golang Template"

func main() {
	log := utils.NewLogger(10)
	log.Add(TAG, MESSAGE)
	log.Add(TAG, CHANGELOG)
	config, err := config.NewConfig("config/config.yaml")
	if err != nil {
		log.Add(TAG, fmt.Sprint("Error reading config file: ", err))
		log.Add(TAG, "Exiting application")
	}
	fmt.Printf("API config: %+v\n", config.SubConfigC)
	// Create a new controller instance
	services.NewController(config)
}
