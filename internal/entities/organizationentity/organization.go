package organizationentity

import "github.com/google/uuid"

type Organization struct {
	ID        uuid.UUID
	OwnerID   uuid.UUID
	CountryID int
	Verified  bool
	Name      string
}
