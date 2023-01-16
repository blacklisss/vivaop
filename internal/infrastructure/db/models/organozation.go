package models

import "github.com/google/uuid"

type Organization struct {
	ID        uuid.UUID
	OwnerID   uuid.UUID
	CountryID int
	Verified  bool
	Name      string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
