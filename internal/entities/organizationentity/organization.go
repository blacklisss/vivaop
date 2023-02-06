package organizationentity

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID               uuid.UUID
	OwnerID          uuid.UUID
	CountryID        int32
	Verified         bool
	Name             string
	RegistrationCode string
	RegistrationDate time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}
