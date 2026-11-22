package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoogleClientID     string
	GoogleClientSecret string
	GoogleCallbackURL  string
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("Erro ao ler arquivo .env: %w", err)
	}

	cfg := Config{
		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		GoogleCallbackURL:  os.Getenv("GOOGLE_CALLBACK_URL"),
	}

	if cfg.GoogleClientID == "" || cfg.GoogleClientSecret == "" || cfg.GoogleCallbackURL == "" {
		return Config{}, fmt.Errorf("Configurações do Google OAuth incompletas")
	}

	return cfg, nil
}
