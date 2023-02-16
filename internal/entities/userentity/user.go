package userentity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID
	CountryID     int32
	FName         string
	MName         string
	LName         string
	Email         string
	Phone         string
	Password      string
	Birthday      string
	VerifiedEmail bool
	Verified      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type EmailVerification struct {
	UserID    uuid.UUID
	Token     string
	ExpiredAt time.Time
	CreatedAt time.Time
}