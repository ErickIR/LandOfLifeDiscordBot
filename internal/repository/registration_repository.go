package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "modernc.org/sqlite"

	"github.com/erickir/LandOfLifeDiscordBot/internal/domain"
)

type sqliteRegistrationRepository struct {
	db *sql.DB
}

func NewRegistrationRepository(db *sql.DB) (*sqliteRegistrationRepository, error) {
	repo := &sqliteRegistrationRepository{db: db}
	if err := repo.ensureSchema(); err != nil {
		return nil, fmt.Errorf("prepare sqlite repository: %w", err)
	}
	return repo, nil
}

var ErrRegistrationAlreadyExists = errors.New("registration already exists")

func (r *sqliteRegistrationRepository) ensureSchema() error {
	const schema = `CREATE TABLE IF NOT EXISTS registrations (
		id TEXT PRIMARY KEY,
		date TEXT NOT NULL,
		event_hour INTEGER NOT NULL,
		channel INTEGER NOT NULL,
		discord_user_id TEXT,
		username TEXT NOT NULL,
		class TEXT,
		level TEXT,
		pet TEXT,
		created_at TEXT NOT NULL
	);`

	if _, err := r.db.Exec(schema); err != nil {
		return err
	}

	const uniqueIndex = `CREATE UNIQUE INDEX IF NOT EXISTS idx_registrations_date_hour_channel_username
	ON registrations(date, event_hour, username);`

	_, err := r.db.Exec(uniqueIndex)
	return err
}

func (r *sqliteRegistrationRepository) Create(ctx context.Context, registration domain.Registration) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO registrations (
			id, date, event_hour, channel, discord_user_id,
			username, class, level, pet, created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		registration.ID,
		registration.Date,
		int(registration.EventHour),
		int(registration.Channel),
		registration.DiscordUserID,
		registration.Username,
		registration.Class,
		registration.Level,
		registration.Pet,
		registration.CreatedAt.Format(time.RFC3339),
	)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return ErrRegistrationAlreadyExists
		}
	}
	return err
}

func (r *sqliteRegistrationRepository) DeleteRegistrationForUser(ctx context.Context, username string, hour domain.Hour, channel domain.Channel) (bool, error) {
	res, err := r.db.ExecContext(ctx,
		`DELETE FROM registrations WHERE username = ? AND event_hour = ? AND channel = ?`,
		username,
		int(hour),
		int(channel),
	)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *sqliteRegistrationRepository) ListByDate(ctx context.Context, regDate string) ([]domain.Registration, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, date, event_hour, channel, discord_user_id, username, class, level, pet, created_at FROM registrations WHERE date = ?`,
		regDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanRows(rows)
}

func (r *sqliteRegistrationRepository) ListByDateAndUser(ctx context.Context, regDate, username string) ([]domain.Registration, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, date, event_hour, channel, discord_user_id, username, class, level, pet, created_at FROM registrations WHERE date = ? AND username = ?`,
		regDate,
		username,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanRows(rows)
}

func (r *sqliteRegistrationRepository) CountByDateHourChannel(ctx context.Context, regDate string, hour domain.Hour, channel domain.Channel) (int, error) {
	var count int
	row := r.db.QueryRowContext(ctx,
		`SELECT COUNT(1) FROM registrations WHERE date = ? AND event_hour = ? AND channel = ?`,
		regDate,
		int(hour),
		int(channel),
	)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *sqliteRegistrationRepository) ExistsByDateHourUser(ctx context.Context, regDate string, hour domain.Hour, username string) (bool, error) {
	var exists int
	row := r.db.QueryRowContext(ctx,
		`SELECT 1 FROM registrations WHERE date = ? AND event_hour = ? AND username = ? LIMIT 1`,
		regDate,
		int(hour),
		username,
	)
	if err := row.Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return exists == 1, nil
}

func (
	r *sqliteRegistrationRepository,
) DeleteOlderThan(
	ctx context.Context,
	regDate string,
) error {
	_, err := r.db.ExecContext(ctx,
		`DELETE FROM registrations WHERE date < ?`,
		regDate,
	)
	return err
}

func (r *sqliteRegistrationRepository) scanRows(rows *sql.Rows) ([]domain.Registration, error) {
	var registrations []domain.Registration
	for rows.Next() {
		registration, err := r.scanRow(rows)
		if err != nil {
			return nil, err
		}
		registrations = append(registrations, registration)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return registrations, nil
}

func (r *sqliteRegistrationRepository) scanRow(rows *sql.Rows) (domain.Registration, error) {
	var registration domain.Registration
	var eventHour int
	var channel int
	var createdAt string

	if err := rows.Scan(
		&registration.ID,
		&registration.Date,
		&eventHour,
		&channel,
		&registration.DiscordUserID,
		&registration.Username,
		&registration.Class,
		&registration.Level,
		&registration.Pet,
		&createdAt,
	); err != nil {
		return registration, err
	}

	registration.EventHour = domain.Hour(eventHour)
	registration.Channel = domain.Channel(channel)
	var err error
	registration.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return registration, err
	}

	return registration, nil
}
