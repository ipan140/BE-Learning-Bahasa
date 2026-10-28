package config

import (
	"log"
	"os"
)

type Config struct {
	Port           string
	DatabaseURL    string
	JWTSecret      string
	GoogleClientID string
	FrontendURL    string
}

var AppConfig *Config

func InitConfig() *Config {
	LoadEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("⚠️ Peringatan: DATABASE_URL tidak ditemukan di file .env")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Println("⚠️ Peringatan: JWT_SECRET tidak ditemukan di file .env")
	}

	AppConfig = &Config{
		Port:           port,
		DatabaseURL:    dbURL,
		JWTSecret:      jwtSecret,
		GoogleClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		FrontendURL:    os.Getenv("FRONTEND_URL"),
	}
	return AppConfig
}
