package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"vivaop/internal/entities/countryentity"
	"vivaop/internal/usecases/app/repos/countryrepo"

	"github.com/pkg/errors"
)

type Country struct {
	ID     int32  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	NameEn string `json:"name_en,omitempty"`
	Code   string `json:"code,omitempty"`
}

func (rt *Handlers) CreateCountry(ctx context.Context, args countryrepo.CreateCountryParams) (Country, error) {
	c, err := rt.cs.CreateCountry(ctx, args)
	if err != nil {
		return Country{}, fmt.Errorf("error when creating: %w", err)
	}

	return Country{
		ID:     c.ID,
		Name:   c.Name,
		NameEn: c.NameEn,
		Code:   c.Code,
	}, nil
}

var ErrCountryNotFound = errors.New("country not found")

func (rt *Handlers) GetCountry(ctx context.Context, id int32) (Country, error) {
	c, err := rt.cs.GetCountry(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Country{}, ErrCountryNotFound
		}
		return Country{}, fmt.Errorf("error when reading: %w", err)
	}

	return Country{
		ID:     c.ID,
		Name:   c.Name,
		NameEn: c.NameEn,
		Code:   c.Code,
	}, nil
}

func (rt *Handlers) DeleteCountry(ctx context.Context, id int32) (Country, error) {
	c, err := rt.cs.DeleteCountry(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Country{}, ErrCountryNotFound
		}
		return Country{}, fmt.Errorf("error when reading: %w", err)
	}

	return Country{
		ID:     c.ID,
		Name:   c.Name,
		NameEn: c.NameEn,
		Code:   c.Code,
	}, nil
}

func (rt *Handlers) ListCountries(ctx context.Context) ([]*countryentity.Country, error) {
	c, err := rt.cs.ListCountries(ctx)
	if err != nil {
		return nil, fmt.Errorf("error when getting: %w", err)
	}

	return c, nil
}
