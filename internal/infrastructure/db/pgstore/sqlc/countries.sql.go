// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: countries.sql

package pgstore

import (
	"context"
	"vivaop/internal/entities/countryentity"
	"vivaop/internal/usecases/app/repos/countryrepo"
)

const createCountry = `-- name: CreateCountry :one
INSERT INTO countries (name, name_en, code)
VALUES ($1, $2, $3)
RETURNING id, name, name_en, code, created_at, updated_at, deleted_at
`

func (q *Queries) CreateCountry(ctx context.Context, arg *countryrepo.CreateCountryParams) (*countryentity.Country, error) {
	row := q.db.QueryRowContext(ctx, createCountry, arg.Name, arg.NameEn, arg.Code)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.NameEn,
		&i.Code,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &countryentity.Country{
		ID:     i.ID,
		Name:   i.Name,
		NameEn: i.NameEn,
		Code:   i.Code,
	}, err
}

const deleteCountry = `-- name: DeleteCountry :exec
DELETE
FROM countries
WHERE id = $1
`

func (q *Queries) DeleteCountry(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCountry, id)
	return err
}

const getCountry = `-- name: GetCountry :one
SELECT id, name, name_en, code, created_at, updated_at, deleted_at
FROM countries
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCountry(ctx context.Context, id int32) (*countryentity.Country, error) {
	row := q.db.QueryRowContext(ctx, getCountry, id)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.NameEn,
		&i.Code,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &countryentity.Country{
		ID:     i.ID,
		Name:   i.Name,
		NameEn: i.NameEn,
		Code:   i.Code,
	}, err
}

const listCountries = `-- name: ListCountries :many
SELECT id, name, name_en, code, created_at, updated_at, deleted_at
FROM countries
`

func (q *Queries) ListCountries(ctx context.Context) ([]*countryentity.Country, error) {
	rows, err := q.db.QueryContext(ctx, listCountries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*countryentity.Country{}
	for rows.Next() {
		var i Country
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.NameEn,
			&i.Code,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		c := &countryentity.Country{
			ID:     i.ID,
			Name:   i.Name,
			NameEn: i.NameEn,
			Code:   i.Code,
		}
		items = append(items, c)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCountry = `-- name: UpdateCountry :one
UPDATE countries
SET name = $2,
    name_en = $3,
    code = $4
WHERE id = $1
RETURNING id, name, name_en, code, created_at, updated_at, deleted_at
`

type UpdateCountryParams struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
	Code   string `json:"code"`
}

func (q *Queries) UpdateCountry(ctx context.Context, arg UpdateCountryParams) (Country, error) {
	row := q.db.QueryRowContext(ctx, updateCountry,
		arg.ID,
		arg.Name,
		arg.NameEn,
		arg.Code,
	)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.NameEn,
		&i.Code,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
