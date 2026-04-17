package core

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/erickir/LandOfLifeDiscordBot/internal/domain"
	"github.com/erickir/LandOfLifeDiscordBot/internal/repository"
)

type RegistrationService struct {
	registrationRepo repository.RegistrationRepository
	nowFn            func() time.Time
}

func NewRegistrationService(registrationRepo repository.RegistrationRepository) *RegistrationService {
	return &RegistrationService{
		registrationRepo: registrationRepo,
		nowFn:            time.Now,
	}
}

func (s *RegistrationService) currentDate() string {
	return s.nowFn().Format("2006-01-02")
}

func (s *RegistrationService) Register(ctx context.Context, registration domain.Registration) error {
	if registration.EventHour == domain.UnknownHour {
		return domain.ErrInvalidHour
	}
	if registration.Channel == domain.UnknownChannel {
		return domain.ErrInvalidChannel
	}

	registration.Date = s.currentDate()
	if registration.CreatedAt.IsZero() {
		registration.CreatedAt = s.nowFn()
	}

	exists, err := s.registrationRepo.ExistsByDateHourUser(ctx, registration.Date, registration.EventHour, registration.Username)
	if err != nil {
		return err
	}
	if exists {
		return domain.ErrAlreadyRegistered
	}

	count, err := s.registrationRepo.CountByDateHourChannel(ctx, registration.Date, registration.EventHour, registration.Channel)
	if err != nil {
		return err
	}
	if count >= domain.MaxUsersPerChannel {
		return domain.ErrSlotFull
	}

	if err := s.registrationRepo.Create(ctx, registration); err != nil {
		return err
	}

	return nil
}

func (s *RegistrationService) Deregister(ctx context.Context, username string, hour domain.Hour, channel domain.Channel) error {
	if hour == domain.UnknownHour {
		return domain.ErrInvalidHour
	}
	if channel == domain.UnknownChannel {
		return domain.ErrInvalidChannel
	}

	deleted, err := s.registrationRepo.DeleteRegistrationForUser(ctx, username, hour, channel)
	if err != nil {
		return err
	}
	if !deleted {
		return domain.ErrRegistrationNotFound
	}
	return nil
}

func (s *RegistrationService) ListRegistrationForSlot(ctx context.Context, date string) ([]domain.Slot, error) {
	if date == "" {
		date = s.currentDate()
	}

	registrations, err := s.registrationRepo.ListByDate(ctx, date)
	if err != nil {
		return nil, err
	}

	slotsMap := make(map[string]domain.Slot)
	for _, hour := range allHours() {
		for _, channel := range allChannels() {
			key := fmt.Sprintf("%d_%d", hour, channel)
			slotsMap[key] = domain.Slot{
				Hour:          hour,
				Channel:       channel,
				Capacity:      domain.MaxUsersPerChannel,
				Registrations: make([]domain.Registration, 0),
			}
		}
	}

	for _, reg := range registrations {
		key := fmt.Sprintf("%d_%d", reg.EventHour, reg.Channel)
		slot, exists := slotsMap[key]
		if !exists {
			continue
		}
		slot.Registrations = append(slot.Registrations, reg)
		slotsMap[key] = slot
	}

	slots := make([]domain.Slot, 0, len(slotsMap))
	for _, slot := range slotsMap {
		slots = append(slots, slot)
	}

	sort.Slice(slots, func(i, j int) bool {
		if slots[i].Hour != slots[j].Hour {
			return slots[i].Hour < slots[j].Hour
		}
		return slots[i].Channel < slots[j].Channel
	})

	return slots, nil
}

func allHours() []domain.Hour {
	return []domain.Hour{
		domain.Hour0100,
		domain.Hour0300,
		domain.Hour0500,
		domain.Hour0700,
		domain.Hour0900,
		domain.Hour1100,
		domain.Hour1300,
		domain.Hour1500,
		domain.Hour1700,
		domain.Hour1900,
		domain.Hour2100,
		domain.Hour2300,
	}
}

func allChannels() []domain.Channel {
	return []domain.Channel{
		domain.Channel1,
		domain.Channel2,
		domain.Channel3,
		domain.Channel4,
		domain.Channel5,
		domain.Channel6,
		domain.Channel7,
	}
}
