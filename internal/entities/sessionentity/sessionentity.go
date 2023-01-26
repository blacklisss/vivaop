package sessionentity

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID //UserID
	RefreshToken string
	UserAgent    string
	ClientIp     string
	IsBlocked    bool
	ExpiresAt    time.Time
}
