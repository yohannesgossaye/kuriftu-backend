package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	DBURL     string
	JWTSecret string
	Port      string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Warn().Err(err).Msg("Error loading .env file")
	}

	return &Config{
		DBURL:     os.Getenv("DB_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		Port:      os.Getenv("PORT"),
	}, nil
}
