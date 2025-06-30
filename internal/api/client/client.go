package client

import (
	"tugas-cli/internal/api"
	"tugas-cli/internal/config"
	"tugas-cli/internal/utils"
	"fmt"
)

type Client struct {
	config *config.Config
	log    *utils.Logger
}

func NewClient(config *config.Config) *Client {
	log := utils.NewLogger(10)
	return &Client{log: log, config: config}
}

func (c *Client) Run() error {
	apiClient := api.NewAPI(c.config)
	weather, err := apiClient.GetWeather()
	if err != nil {
		c.log.Add("CLIENT", fmt.Sprintf("Failed to get weather: %v", err))
		return err
	}

	c.log.Add("CLIENT", fmt.Sprintf("Cuaca di %s: %.1f C", weather.Location.Name, weather.Current.TempC))
	return nil
}
