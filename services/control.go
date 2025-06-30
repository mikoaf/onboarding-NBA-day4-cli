package services

import (
	"tugas-cli/internal/api/client"
	"tugas-cli/internal/api/server"
	"tugas-cli/internal/config"
	"tugas-cli/internal/hardware"
	"tugas-cli/internal/models"
	"tugas-cli/internal/utils"
	"fmt"
)

const TAG = "SERVICE"

var (
	Config      *config.Config
	Transaction *models.Transaction
	Hardware    *hardware.Hardware
	Client      *client.Client
	Server      *server.Server
	log         *utils.Logger
)

func NewController(config *config.Config) {
	// Starting Message
	log = utils.NewLogger(10)
	log.Add(TAG, "Tugas-CLI")

	// Creating all services instance here
	// transaction := models.NewTransaction()
	// hardware := hardware.NewHardware(config)
	Client := client.NewClient(config)
	// go Client.Run()
	Server := server.NewServer(config)
	// go Server.Run()

	go func(){
		if err := Client.Run(); err != nil{
			fmt.Printf("Client error: %v\n", err)
		}
	}()

	go func(){
		if err := Server.Run(); err != nil{
			fmt.Printf("Server error: %v\n", err)
		}
	}()
	
	select {}
}
