package core

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	_ "modernc.org/sqlite"

	"github.com/erickir/LandOfLifeDiscordBot/internal/domain"
	"github.com/erickir/LandOfLifeDiscordBot/internal/repository"
)

func TestRegistrationService_Register_ReturnsAlreadyRegisteredForDuplicateSlot(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("open sqlite db: %v", err)
	}
	defer db.Close()

	repo, err := repository.NewRegistrationRepository(db)
	if err != nil {
		t.Fatalf("create repo: %v", err)
	}

	service := NewRegistrationService(repo)
	ctx := context.Background()

	registration := domain.Registration{
		ID:            "id-1",
		EventHour:     domain.Hour0100,
		Channel:       domain.Channel1,
		DiscordUserID: "discord-1",
		Username:      "test-user",
		Class:         "Warrior",
		Level:         "99",
		Pet:           "Dragon",
		CreatedAt:     time.Now(),
	}

	if err := service.Register(ctx, registration); err != nil {
		t.Fatalf("first register failed: %v", err)
	}

	if err := service.Register(ctx, registration); err == nil {
		t.Fatal("expected duplicate registration to fail, got nil")
	} else if !errors.Is(err, domain.ErrAlreadyRegistered) {
		t.Fatalf("expected ErrAlreadyRegistered, got %v", err)
	}
}
