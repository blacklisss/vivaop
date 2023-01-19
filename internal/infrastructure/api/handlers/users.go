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

type User struct {
	ID        uuid.UUID
	CountryID int32
	FName     string
	MName     string
	LName     string
	Email     string
	Phone     string
	Password  string
	Birthday  string
}

func (rt *Handlers) CreateUser(ctx context.Context, args *userrepo.CreateUserParams) (*userentity.User, error) {
	user, err := rt.us.CreateUser(ctx, args)
	if err != nil {
		return &userentity.User{}, fmt.Errorf("error when creating: %w", err)
	}

	fmt.Println(user)
	return &userentity.User{
		ID:        user.ID,
		FName:     user.FName,
		MName:     user.MName,
		LName:     user.LName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
		Birthday:  user.Birthday,
		CountryID: user.CountryID,
		CreatedAt: user.CreatedAt,
	}, nil
}

var ErrUserNotFound = errors.New("user not found")

func (rt *Handlers) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	user, err := rt.us.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrUserNotFound
		}
		return User{}, fmt.Errorf("error when reading: %w", err)
	}

	return User{
		ID:       user.ID,
		FName:    user.FName,
		MName:    user.MName,
		LName:    user.LName,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
		Birthday: user.Birthday,
	}, nil
}

func (rt *Handlers) DeleteUser(ctx context.Context, id uuid.UUID) (User, error) {
	user, err := rt.us.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrUserNotFound
		}
		return User{}, fmt.Errorf("error when delete: %w", err)
	}

	return User{
		ID:       user.ID,
		FName:    user.FName,
		MName:    user.MName,
		LName:    user.LName,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
		Birthday: user.Birthday,
	}, nil
}
