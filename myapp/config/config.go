package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	VaultAddr   string `yaml:"vault_addr"`
	AppRoleID   string `yaml:"app_role_id"`
	AppSecretID string `yaml:"app_secret_id"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		return nil, err
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		log.Printf("Error unmarshaling config file: %v", err)
		return nil, err
	}

	return cfg, nil
}
