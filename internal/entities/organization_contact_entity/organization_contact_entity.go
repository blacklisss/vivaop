package organization_contact_entity

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationContact struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Name           string
	Phone          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
