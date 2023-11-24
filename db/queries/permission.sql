-- name: CreatePermission :one
INSERT INTO permissions (name)
VALUES ($1)
RETURNING *;

-- name: GetPermissionByID :one
SELECT *
FROM permissions
WHERE id = $1;

-- name: GetAllPermissions :many
SELECT *
FROM permissions;

-- name: UpdatePermissionByID :one
UPDATE permissions
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeletePermissionByID :one
DELETE
FROM permissions
WHERE id = $1
RETURNING *;
