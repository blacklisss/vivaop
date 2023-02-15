-- name: CreateOrganizationContact :one
INSERT INTO organization_contacts
(id,
 organization_id,
 name,
 phone
)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOrganizationContact :one
SELECT *
FROM organization_contacts
WHERE id = $1
LIMIT 1;

-- name: ListOrganizationContacts :many
SELECT *
FROM organization_contacts
WHERE organization_id = $1
  AND deleted_at IS NULL;

-- name: DeleteOrganizationContact :one
UPDATE organization_contacts
SET deleted_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING *;