package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"URLs.log"`
}

var (
	serverAddress   string
	baseURL         string
	fileStoragePath string
)

func init() {
	flag.StringVar(&serverAddress, "a", "", "server address")
	flag.StringVar(&baseURL, "b", "", "base URL")
	flag.StringVar(&fileStoragePath, "f", "", "file storage path")
}

func NewConfig() *Config {
	flag.Parse()
	log.Printf("server address: %s, base URL: %s, file storagePath: %s\n", serverAddress, baseURL, fileStoragePath)

	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	if serverAddress != "" {
		cfg.ServerAddress = serverAddress
	}
	if baseURL != "" {
		cfg.BaseURL = baseURL
	}
	if fileStoragePath != "" {
		cfg.FileStoragePath = fileStoragePath
	}

	return &cfg
}
