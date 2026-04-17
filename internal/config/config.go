package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env          string
	TokenID      string
	GuildID      string
	DatabasePath string
}

func NewConfig() *Config {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// Load .env file if it exists, but ignore error if it doesn't exist
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Error loading .env file: %v", err)
	}

	tokenID := os.Getenv("BOT_TOKEN")
	guildID := os.Getenv("GUILD_ID")
	databasePath := os.Getenv("DATABASE_PATH")
	if databasePath == "" {
		databasePath = "registrations.db"
	}

	return &Config{
		Env:          env,
		TokenID:      tokenID,
		GuildID:      guildID,
		DatabasePath: databasePath,
	}
}
