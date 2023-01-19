-- name: CreateUser :one
INSERT INTO users (
    id,
    fname,
    mname,
    lname,
    email,
    phone,
    password,
    birthdate,
    country_id
) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9 ) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByPhone :one
SELECT * FROM users
WHERE phone = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;
