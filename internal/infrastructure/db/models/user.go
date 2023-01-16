package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	CountryID int
	FName     string
	MName     string
	LName     string
	Email     string
	Phone     string
	Password  string
	Birthday  string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
