package sessionrepo

import (
	"context"
	"fmt"
	"time"
	"vivaop/internal/entities/sessionentity"

	"github.com/google/uuid"
)

type CreateSessionParams struct {
	ID           uuid.UUID `json:"id"` // userID
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type SessionStore interface {
	CreateSession(ctx context.Context, args *CreateSessionParams) (*sessionentity.Session, error)
	GetSession(ctx context.Context, userID uuid.UUID) (*sessionentity.Session, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type Session struct {
	sstore SessionStore
}

func (s *Session) CreateSession(ctx context.Context, args *CreateSessionParams) (*sessionentity.Session, error) {
	session, err := s.sstore.CreateSession(ctx, args)
	if err != nil {
		return nil, fmt.Errorf("create session error: %w", err)
	}

	return session, nil
}

func (s *Session) GetSession(ctx context.Context, id uuid.UUID) (*sessionentity.Session, error) {
	session, err := s.sstore.GetSession(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get session error: %w", err)
	}

	return session, nil
}

func NewSession(sstore SessionStore) *Session {
	return &Session{
		sstore,
	}
}
