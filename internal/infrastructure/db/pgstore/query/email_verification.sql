-- name: CreateEmailVerification :one
INSERT INTO verify_email (user_id,
                          token,
                          expired_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetEmailVerification :one
SELECT * FROM verify_email
WHERE token = $1 LIMIT 1;

-- name: DeleteEmailVerification :exec
DELETE
FROM  verify_email
WHERE user_id = $1;