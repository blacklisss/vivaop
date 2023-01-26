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

func (rt *Handlers) CreateUser(ctx context.Context, args *userrepo.CreateUserParams) (*userentity.User, error) {
	user, err := rt.us.CreateUser(ctx, args)
	if err != nil {
		return &userentity.User{}, fmt.Errorf("error when creating: %w", err)
	}

	return user, nil
}

var ErrUserNotFound = errors.New("user not found")

func (rt *Handlers) GetUserByID(ctx context.Context, id uuid.UUID) (*userentity.User, error) {
	user, err := rt.us.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &userentity.User{}, ErrUserNotFound
		}
		return &userentity.User{}, fmt.Errorf("error when reading: %w", err)
	}

	return user, nil
}

func (rt *Handlers) GetUserByEmail(ctx context.Context, email string) (*userentity.User, error) {
	user, err := rt.us.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &userentity.User{}, ErrUserNotFound
		}
		return &userentity.User{}, fmt.Errorf("error when reading: %w", err)
	}

	return user, nil
}

func (rt *Handlers) DeleteUser(ctx context.Context, id uuid.UUID) (*userentity.User, error) {
	user, err := rt.us.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &userentity.User{}, ErrUserNotFound
		}
		return &userentity.User{}, fmt.Errorf("error when delete: %w", err)
	}

	return user, nil
}
