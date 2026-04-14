package domain

import "errors"

var (
	ErrInvalidHour          = errors.New("invalid hour")
	ErrInvalidChannel       = errors.New("invalid channel")
	ErrSlotFull             = errors.New("slot is full")
	ErrAlreadyRegistered    = errors.New("user already registered for this hour")
	ErrRegistrationNotFound = errors.New("registration not found")
)
