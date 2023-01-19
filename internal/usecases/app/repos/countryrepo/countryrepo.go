package countryrepo

import (
	"context"
	"fmt"
	"vivaop/internal/entities/countryentity"
)

type CreateCountryParams struct {
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
	Code   string `json:"code"`
}

type CountryStore interface {
	CreateCountry(ctx context.Context, arg CreateCountryParams) (*countryentity.Country, error)
	GetCountry(ctx context.Context, id int32) (*countryentity.Country, error)
	ListCountries(ctx context.Context) ([]*countryentity.Country, error)
	DeleteCountry(ctx context.Context, id int32) error
}

type Countries struct {
	cstore CountryStore
}

func NewCountries(cstore CountryStore) *Countries {
	return &Countries{
		cstore,
	}
}

func (cs *Countries) CreateCountry(ctx context.Context, arg CreateCountryParams) (*countryentity.Country, error) {
	country, err := cs.cstore.CreateCountry(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("create country error: %w", err)
	}
	return country, nil
}

func (cs *Countries) GetCountry(ctx context.Context, id int32) (*countryentity.Country, error) {
	country, err := cs.cstore.GetCountry(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("read country error: %w", err)
	}
	return country, nil
}

func (cs *Countries) ListCountries(ctx context.Context) ([]*countryentity.Country, error) {
	countries, err := cs.cstore.ListCountries(ctx)
	if err != nil {
		return nil, fmt.Errorf("read country error: %w", err)
	}
	return countries, nil
}

func (cs *Countries) DeleteCountry(ctx context.Context, id int32) (*countryentity.Country, error) {
	country, err := cs.cstore.GetCountry(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("delete country error: %w", err)
	}
	return country, cs.cstore.DeleteCountry(ctx, id)
}
