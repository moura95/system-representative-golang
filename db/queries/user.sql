-- name: AddUserPermission :one
INSERT INTO user_permissions (user_id, permission_id)
VALUES ($1, $2)
RETURNING *;

-- name: CreateUser :one
INSERT INTO users (representative_id, cpf, first_name, last_name, email, password, phone)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: DeleteUserByID :one
UPDATE users
SET is_active = FALSE
WHERE id = $1
  AND is_active = TRUE
RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserPasswordByID :one
SELECT password
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: ListUsersByRepresentativeID :many
SELECT *
FROM users
WHERE representative_id = $1
  AND is_active = $2;

-- name: RemoveUserPermissionByID :one
DELETE
FROM user_permissions
WHERE user_id = $1
  AND permission_id = $2
RETURNING *;

-- name: RemoveUserByID :one
DELETE
FROM users
WHERE id = $1
RETURNING *;

-- name: RestoreUserByID :one
UPDATE users
SET is_active  = TRUE,
    updated_at = NOW()
WHERE id = $1
  AND is_active = FALSE
RETURNING *;

-- name: UpdateLastLogin :exec
UPDATE users
SET last_login = NOW()
WHERE id = $1;

-- name: UpdateUserByID :one
UPDATE users
SET password   = COALESCE(sqlc.narg('password'), password),
    first_name = COALESCE(sqlc.narg('first_name'), first_name),
    last_name  = COALESCE(sqlc.narg('last_name'), last_name),
    phone      = COALESCE(sqlc.narg('phone'), phone),
    cpf        = COALESCE(sqlc.narg('cpf'), cpf),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: ChangePasswordUserByID :exec
UPDATE users
SET password   = sqlc.arg('new_password'),
    updated_at = NOW()
WHERE id = $1;

-- name: GetUserPermissionAndName :many
SELECT p.name
FROM user_permissions u
         JOIN permissions p
              ON u.permission_id = p.id
WHERE u.user_id = $1;

-- name: GetEmailAndNameByRepresentativeID :one
SELECT email, name
FROM representatives
WHERE id = $1;
