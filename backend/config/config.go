package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	DatabaseURL       string
	JWTSecret         string
	HasuraAdminSecret string
	HasuraEndpoint    string
	CloudinaryName    string
	CloudinaryAPIKey  string
	CloudinarySecret  string
	ChapaSecretKey    string
	ChapaPublicKey    string
	ChapaCallbackURL  string
	ChapaReturnURL    string
}

func Load() *Config {
	godotenv.Load()

	cfg := &Config{
		Port:              getEnv("PORT", "3001"),
		DatabaseURL:       getEnv("DATABASE_URL", ""),
		JWTSecret:         mustGetEnv("JWT_SECRET"),
		HasuraAdminSecret: mustGetEnv("HASURA_ADMIN_SECRET"),
		HasuraEndpoint:    mustGetEnv("HASURA_ENDPOINT"),
		CloudinaryName:    getEnv("CLOUDINARY_CLOUD_NAME", ""),
		CloudinaryAPIKey:  getEnv("CLOUDINARY_API_KEY", ""),
		CloudinarySecret:  getEnv("CLOUDINARY_API_SECRET", ""),
		ChapaSecretKey:    getEnv("CHAPA_SECRET_KEY", ""),
		ChapaPublicKey:    getEnv("CHAPA_PUBLIC_KEY", ""),
		ChapaCallbackURL:  getEnv("CHAPA_CALLBACK_URL", ""),
		ChapaReturnURL:    getEnv("CHAPA_RETURN_URL", ""),
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable is required", key)
	}
	return value
}
