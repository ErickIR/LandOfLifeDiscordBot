package config

import (
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("ENV")
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
