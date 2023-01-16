package countryrepo

import (
	"context"
	"net/http"
	"vivaop/internal/entities/countryentity"
)

type CountryStore interface {
	Create(ctx context.Context, l string) (*countryentity.Country, error)
	Read(ctx context.Context, hash string, r *http.Request) (*countryentity.Country, error)
	Delete(ctx context.Context, hash string, r *http.Request) error
}

type Countries struct {
	cstore CountryStore
}

func NewCountries(cstore CountryStore) *Countries {
	return &Countries{
		cstore,
	}
}
