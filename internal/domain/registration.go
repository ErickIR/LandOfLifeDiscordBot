package domain

import "time"

type Registration struct {
	ID            string
	Date          string
	EventHour     Hour
	Channel       Channel
	DiscordUserID string
	Username      string
	Class         string
	Level         string
	Pet           string
	CreatedAt     time.Time
}

type Slot struct {
	Hour          Hour
	Channel       Channel
	Capacity      int
	Registrations []Registration
}
