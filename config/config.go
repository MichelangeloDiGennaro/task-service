package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AWS struct {
		Endpoint string `yaml:"endpoint"`
		Region   string `yaml:"region"`
	} `yaml:"aws"`
}

func LoadConfig() Config {
	env := os.Getenv("APP_ENV")
	configFile := "environments/local.yaml"
	if env == "production" {
		configFile = "environments/prod.yaml"
	}

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Errore nella lettura del file di configurazione: %v", err)
	}
	defer file.Close()

	configData, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Errore nella lettura del file di configurazione: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Errore nel parsing del file di configurazione: %v", err)
	}

	return config
}
