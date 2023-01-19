package userentity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	CountryID int32
	FName     string
	MName     string
	LName     string
	Email     string
	Phone     string
	Password  string
	Birthday  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
