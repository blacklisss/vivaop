package userrepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
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
	DeleteEmailVerification(ctx context.Context, userID uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	CreateEmailVerification(ctx context.Context, args *CreateEmailVerificationParams) (*userentity.EmailVerification, error)
	GetEmailVerification(ctx context.Context, token string) (*userentity.EmailVerification, error)
	VerifyUserEmail(ctx context.Context, id uuid.UUID) (*userentity.User, error)
}

type Users struct {
	ustore UserStore
}

func NewUsers(ustore UserStore) *Users {
	return &Users{
		ustore,
	}
}

type CreateEmailVerificationParams struct {
	UserID    uuid.UUID `json:"user_id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
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

func (us *Users) CreateEmailVerification(ctx context.Context, args *CreateEmailVerificationParams) (*userentity.EmailVerification, error) {
	emailverification, err := us.ustore.CreateEmailVerification(ctx, args)
	if err != nil {
		return nil, fmt.Errorf("cerate email verification error: %w", err)
	}

	return emailverification, nil
}

func (us *Users) GetEmailVerification(ctx context.Context, token string) (*userentity.EmailVerification, error) {
	emailVerification, err := us.ustore.GetEmailVerification(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("get email verification error: %w", err)
	}

	return emailVerification, nil
}

func (us *Users) VerifyUserEmail(ctx context.Context, id uuid.UUID) (*userentity.User, error) {
	user, err := us.ustore.VerifyUserEmail(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("email verification error: %w", err)
	}

	return user, nil
}

func (us *Users) DeleteEmailVerification(ctx context.Context, userID uuid.UUID) error {
	err := us.ustore.DeleteEmailVerification(ctx, userID)
	if err != nil {
		return fmt.Errorf("email verification delete error: %w", err)
	}

	return nil
}
