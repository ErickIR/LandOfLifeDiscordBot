package repository

import (
	"context"

	"github.com/erickir/LandOfLifeDiscordBot/internal/domain"
)

type RegistrationRepository interface {
	Create(ctx context.Context, registration domain.Registration) error
	DeleteRegistrationForUser(ctx context.Context, username string, hour domain.Hour, channel domain.Channel) (bool, error)
	ListByDate(ctx context.Context, regDate string) ([]domain.Registration, error)
	ListByDateAndUser(ctx context.Context, regDate, username string) ([]domain.Registration, error)
	CountByDateHourChannel(ctx context.Context, regDate string, hour domain.Hour, channel domain.Channel) (int, error)
	ExistsByDateHourUser(ctx context.Context, regDate string, hour domain.Hour, username string) (bool, error)
	DeleteOlderThan(ctx context.Context, regDate string) error
}
