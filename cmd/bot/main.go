package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/erickir/LandOfLifeDiscordBot/internal/bot"
	"github.com/erickir/LandOfLifeDiscordBot/internal/commands"
	"github.com/erickir/LandOfLifeDiscordBot/internal/config"
	"github.com/erickir/LandOfLifeDiscordBot/internal/core"
	"github.com/erickir/LandOfLifeDiscordBot/internal/repository"
	"github.com/erickir/LandOfLifeDiscordBot/pkg/discord"
)

func main() {
	configuration := config.NewConfig()

	db, err := sql.Open("sqlite", configuration.DatabasePath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if err := db.PingContext(context.Background()); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	registrationRepo, err := repository.NewRegistrationRepository(db)
	if err != nil {
		log.Fatalf("Failed to initialize registration repository: %v", err)
	}

	registrationService := core.NewRegistrationService(registrationRepo)

	router, err := bot.NewRouter(
		commands.NewLandOfLifeCommand(registrationService),
	)
	if err != nil {
		log.Fatalf("Failed to create bot client: %v", err)
	}

	client, err := discord.NewClient(configuration, router)
	if err != nil {
		log.Fatalf("Failed to create Discord client: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := client.Open(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Signal received, shutting down bot...")
	if err := client.Close(); err != nil {
		log.Printf("Error closing bot: %v", err)
	}
	log.Println("Bot stopped.")
}
