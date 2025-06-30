package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// This file contains all the application configurations
// Save all configuration struct in one place and just derived later by its main struct

type Config struct {
	SubConfigA SubConfigA `yaml:"sub_config_a"`
	SubConfigB SubConfigB `yaml:"sub_config_b"`
	SubConfigC SubConfigC `yaml:"sub_config_c"`
}

type SubConfigA struct {
	FieldA string  `yaml:"field_a"`
	FieldB int64   `yaml:"field_b"`
	FieldC float64 `yaml:"field_c"`
	FieldD bool    `yaml:"field_d"`
}

type SubConfigB struct {
	FieldA string  `yaml:"field_a"`
	FieldB int64   `yaml:"field_b"`
	FieldC float64 `yaml:"field_c"`
	FieldD bool    `yaml:"field_d"`
}

type SubConfigC  struct {
	ApiUrl string `yaml:"api_url"`
	ApiKey string `yaml:"api_key"`
	City string `yaml:"default_city"`
}

// Read config file
func NewConfig(filePath string) (cfg *Config, err error) {
	// Allocates mememory for configuration data
	config := new(Config)

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("reading file %s error: %v", filePath, err)
		return nil, err
	}

	// Unmarshal the file and save to config memory
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Printf("parsing file %s error: %v", filePath, err)
		return nil, err
	}

	return config, nil
}
