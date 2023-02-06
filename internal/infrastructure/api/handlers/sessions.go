package handlers

import (
	"context"
	"fmt"
	"vivaop/internal/entities/sessionentity"
	"vivaop/internal/usecases/app/repos/sessionrepo"

	"github.com/google/uuid"
)

func (rt *Handlers) CreateSession(ctx context.Context, args *sessionrepo.CreateSessionParams) (*sessionentity.Session, error) {
	session, err := rt.ss.CreateSession(ctx, args)
	if err != nil {
		return nil, fmt.Errorf("create session error: %w", err)
	}

	return session, nil
}

func (rt *Handlers) GetSession(ctx context.Context, id uuid.UUID) (*sessionentity.Session, error) {
	session, err := rt.ss.GetSession(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get session error: %w", err)
	}

	return session, nil
}

func (rt *Handlers) DeleteSession(ctx context.Context, id uuid.UUID) error {
	return rt.ss.DeleteSession(ctx, id)
}
