package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	_ "modernc.org/sqlite"

	"github.com/erickir/LandOfLifeDiscordBot/internal/domain"
)

func TestRegistrationRepository_Create_DuplicateRegistrationFails(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("open sqlite db: %v", err)
	}
	defer db.Close()

	repo, err := NewRegistrationRepository(db)
	if err != nil {
		t.Fatalf("create repo: %v", err)
	}

	ctx := context.Background()
	registration := domain.Registration{
		ID:        "id-1",
		Date:      "2026-04-17",
		EventHour: domain.Hour0100,
		Channel:   domain.Channel1,
		Username:  "test-user",
		CreatedAt: time.Now(),
	}

	if err := repo.Create(ctx, registration); err != nil {
		t.Fatalf("first insert failed: %v", err)
	}

	duplicate := registration
	duplicate.ID = "id-2"

	if err := repo.Create(ctx, duplicate); err == nil {
		t.Fatal("expected duplicate insert to fail, got nil")
	} else if !errors.Is(err, ErrRegistrationAlreadyExists) {
		t.Fatalf("expected ErrRegistrationAlreadyExists, got %v", err)
	}
}
