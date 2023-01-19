package userrepo

import (
	"context"
	"database/sql"
	"fmt"
	"vivaop/internal/entities/userentity"

	"github.com/google/uuid"
)

type CreateUserParams struct {
	ID        uuid.UUID      `json:"id"`
	Fname     sql.NullString `json:"fname"`
	Mname     sql.NullString `json:"mname"`
	Lname     sql.NullString `json:"lname"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Password  string         `json:"password"`
	Birthdate sql.NullTime   `json:"birthdate"`
	CountryID sql.NullInt32  `json:"country_id"`
}

type UserStore interface {
	CreateUser(ctx context.Context, args *CreateUserParams) (*userentity.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*userentity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*userentity.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*userentity.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type Users struct {
	ustore UserStore
}

func NewUsers(ustore UserStore) *Users {
	return &Users{
		ustore,
	}
}

func (us *Users) CreateUser(ctx context.Context, args *CreateUserParams) (*userentity.User, error) {
	user, err := us.ustore.CreateUser(ctx, args)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}

	return user, nil
}

func (us *Users) GetUserByID(ctx context.Context, id uuid.UUID) (*userentity.User, error) {
	user, err := us.ustore.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}

	return user, nil
}

func (us *Users) GetUserByEmail(ctx context.Context, email string) (*userentity.User, error) {
	user, err := us.ustore.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}

	return user, nil
}

func (us *Users) GetUserByPhone(ctx context.Context, phone string) (*userentity.User, error) {
	user, err := us.ustore.GetUserByEmail(ctx, phone)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}

	return user, nil
}

func (us *Users) DeleteUser(ctx context.Context, id uuid.UUID) (*userentity.User, error) {
	user, err := us.ustore.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("delete country error: %w", err)
	}
	return user, us.ustore.DeleteUser(ctx, id)
}
