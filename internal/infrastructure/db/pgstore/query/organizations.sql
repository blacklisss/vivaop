-- name: CreateOrganization :one
INSERT INTO organizations
(id,
 name,
 country_id,
 owner_id,
 verified,
 registration_code,
 registration_date)
VALUES ($1, $2, $3, $4, $5, $6, $7)
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
WHERE owner_id = $1
  AND deleted_at IS NULL;

-- name: UpdateOrganization :one
UPDATE organizations
SET name              = $2,
    country_id        = $3,
    owner_id          = $4,
    verified          = $5,
    registration_code = $6,
    registration_date = $7,
    updated_at        = NOW()
WHERE id = $1
RETURNING *;

-- name: VerifyOrganization :one
UPDATE organizations
SET verified   = true,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteOrganization :one
UPDATE organizations
SET deleted_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING *;

-- name: AddRegistrationImage :one
UPDATE organizations
SET registrtion_image = $1,
    updated_at        = NOW()
WHERE id = $1
RETURNING *;
