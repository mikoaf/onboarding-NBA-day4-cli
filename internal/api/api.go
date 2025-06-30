package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"

	"tugas-cli/internal/config"
)

type API struct {
	APIKey string
	BaseURL string
	City string
}

func NewAPI(cfg *config.Config) *API {
	return &API{
		APIKey: cfg.SubConfigC.ApiKey,
		BaseURL: cfg.SubConfigC.ApiUrl,
		City: cfg.SubConfigC.City,
	}
}

type WeatherResponse struct{
	Location struct{
		Name string `json:"name"`
	} `json:"location"`
	Current struct{
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func (a *API) GetWeather() (*WeatherResponse, error){
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", a.BaseURL, a.APIKey, a.City)

	resp, err := http.Get(url)
	if err != nil{
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	body,_:= io.ReadAll(resp.Body)

	if closeErr := resp.Body.Close(); closeErr != nil{
		return nil, fmt.Errorf("failed to closed response body: %w", closeErr)
	}

	var data WeatherResponse
	err = json.Unmarshal(body, &data)
	if err != nil{
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}
	return &data, nil
}