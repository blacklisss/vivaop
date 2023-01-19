package countryrepo

import (
	"context"
	"fmt"
	"net/http"
	"vivaop/internal/entities/countryentity"
)

type CountryStore interface {
	Create(ctx context.Context, name string, name_en string, code string) (*countryentity.Country, error)
	Read(ctx context.Context, id int64, r *http.Request) (*countryentity.Country, error)
	ReadAll(ctx context.Context, r *http.Request) ([]*countryentity.Country, error)
	Delete(ctx context.Context, id int64, r *http.Request) error
}

type Countries struct {
	cstore CountryStore
}

func NewCountries(cstore CountryStore) *Countries {
	return &Countries{
		cstore,
	}
}

func (cs *Countries) Create(ctx context.Context, name string, name_en string, code string) (*countryentity.Country, error) {
	country, err := cs.cstore.Create(ctx, name, name_en, code)
	if err != nil {
		return nil, fmt.Errorf("create country error: %w", err)
	}
	return country, nil
}

func (cs *Countries) Read(ctx context.Context, id int64, r *http.Request) (*countryentity.Country, error) {
	country, err := cs.cstore.Read(ctx, id, r)
	if err != nil {
		return nil, fmt.Errorf("read country error: %w", err)
	}
	return country, nil
}

func (cs *Countries) ReadAll(ctx context.Context, r *http.Request) ([]*countryentity.Country, error) {
	countries, err := cs.cstore.ReadAll(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("read country error: %w", err)
	}
	return countries, nil
}

func (cs *Countries) Delete(ctx context.Context, id int64, r *http.Request) (*countryentity.Country, error) {
	country, err := cs.cstore.Read(ctx, id, r)
	if err != nil {
		return nil, fmt.Errorf("delete country error: %w", err)
	}
	return country, cs.cstore.Delete(ctx, id, r)
}
