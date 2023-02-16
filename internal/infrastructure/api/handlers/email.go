package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"vivaop/internal/entities/userentity"
	"vivaop/internal/usecases/app/repos/userrepo"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (rt *Handlers) CreateEmailVerification(ctx context.Context, args *userrepo.CreateEmailVerificationParams) (*userentity.EmailVerification, error) {
	emailVerification, err := rt.us.CreateEmailVerification(ctx, args)
	if err != nil {
		return nil, fmt.Errorf("error when creating email verification: %w", err)
	}

	return emailVerification, nil
}

func (rt *Handlers) CheckEmailToken(ctx context.Context, emailToken string) (*userentity.EmailVerification, error) {
	return rt.us.GetEmailVerification(ctx, emailToken)
}

func (rt *Handlers) VerifyUserEmail(ctx context.Context, id uuid.UUID) (*userentity.User, error) {
	user, err := rt.us.VerifyUserEmail(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &userentity.User{}, ErrUserNotFound
		}
		return &userentity.User{}, fmt.Errorf("error when delete: %w", err)
	}

	return user, nil
}

func (rt *Handlers) DeleteUserEmailVerification(ctx context.Context, id uuid.UUID) error {
	err := rt.us.DeleteEmailVerification(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
