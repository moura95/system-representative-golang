-- name: CreateSmtp :one
INSERT INTO smtp (representative_id, email, password, server, port)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSmtpByRepresentativeID :one
SELECT *
FROM smtp
WHERE representative_id = $1;

-- name: UpdateSmtpByID :one
UPDATE smtp
SET email        = COALESCE(sqlc.narg('email'), email),
    password     = COALESCE(sqlc.narg('password'), password),
    server       = COALESCE(sqlc.narg('server'), server),
    port         = COALESCE(sqlc.narg('port'), port),
    updated_at   = NOW()
WHERE representative_id = $1
RETURNING *;

-- name: DeleteSmtpByRepresentativeID :one
DELETE
FROM smtp
WHERE representative_id = $1
RETURNING *;
