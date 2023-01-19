-- name: CreateCountry :one
INSERT INTO countries (name, name_en, code)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCountry :one
SELECT *
FROM countries
WHERE id = $1
LIMIT 1;

-- name: ListCountries :many
SELECT *
FROM countries;

-- name: UpdateCountry :one
UPDATE countries
SET name = $2,
    name_en = $3,
    code = $4
WHERE id = $1
RETURNING *;

-- name: DeleteCountry :exec
DELETE
FROM countries
WHERE id = $1;
