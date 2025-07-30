package config

import (
	"log"

	"gin/utils"
	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server ServerConfig
}

type DatabaseConfig struct {
	User 		string
	Password 	string
	Host 		string
	Port 		string
	Name 		string
}

type ServerConfig struct {
	Port string
	Mode string // development, production, test
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env found, using environtment variables")
	}

	return &Config{
		Database: DatabaseConfig{
			User: utils.GetEnv("DB_USER", "root"),
			Password: utils.GetEnv("DB_PASSWORD", ""),
			Host: utils.GetEnv("DB_HOST", "localhost"),
			Port: utils.GetEnv("DB_PORT", "3306"),
			Name: utils.GetEnv("DB_NAME", "ayo_league"),
		},
		Server: ServerConfig{
			Port: utils.GetEnv("SERVER_PORT", "8080"),
			Mode: utils.GetEnv("SERVER_MODE", "development"),
		},
	}
}
