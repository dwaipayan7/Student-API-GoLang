package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`

	HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {
	// Load config from file or environment variables
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flag.StringVar(&configPath, "config", "", "path to the configuration file")
		flag.Parse()

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	// Validate that the config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	var cfg Config

	// Read configuration using cleanenv
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read config file: %s", err.Error())
	}

	return &cfg
}
