-- name: CreateOrganization :one
INSERT INTO organizations
(id,
 name,
 country_id,
 owner_id,
 verified)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetOrganization :one
SELECT *
FROM organizations
WHERE id = $1
LIMIT 1;

-- name: GetOrganizationByOwner :one
SELECT *
FROM organizations
WHERE id = $1
  AND owner_id = $2
LIMIT 1;

-- name: ListOwnerOrganization :many
SELECT *
FROM organizations
WHERE owner_id = $1;

-- name: UpdateOrganization :one
UPDATE organizations
SET name       = $2,
    country_id = $3,
    owner_id   = $4,
    verified   = $5
WHERE id = $1
RETURNING *;

-- name: DeleteOrganization :exec
DELETE
FROM organizations
WHERE id = $1;