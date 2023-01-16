package userrepo

import (
	"context"
	"net/http"
	"vivaop/internal/entities/userentity"
)

type UserStore interface {
	Create(ctx context.Context, l string) (*userentity.User, error)
	Read(ctx context.Context, hash string, r *http.Request) (*userentity.User, error)
	Delete(ctx context.Context, hash string, r *http.Request) error
}

type Users struct {
	ustore UserStore
}

func NewUsers(ustore UserStore) *Users {
	return &Users{
		ustore,
	}
}
