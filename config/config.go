package config

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
	AppConfig = &Config{
		Port:           GetEnv("PORT", "8080"),
		DatabaseURL:    GetEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/learning_bahasa?sslmode=disable"),
		JWTSecret:      GetEnv("JWT_SECRET", "super-secret-jwt-key-learning-bahasa-2026"),
		GoogleClientID: GetEnv("GOOGLE_CLIENT_ID", ""),
		FrontendURL:    GetEnv("FRONTEND_URL", "http://localhost:4321"),
	}
	return AppConfig
}
