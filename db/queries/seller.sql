-- name: CreateSellers :one
INSERT INTO sellers (representative_id, name, pix, email, phone, observation, cpf)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: DeleteSellerByID :one
UPDATE sellers
SET is_active = FALSE
WHERE id = $1
  AND is_active = TRUE
RETURNING *;

-- name: GetSellerByID :one
SELECT *
FROM sellers
WHERE id = $1;

-- name: ListSellersByRepresentativeID :many
SELECT *
FROM sellers
WHERE representative_id = $1
AND is_active = $2;

-- name: RemoveSellerByID :one
DELETE
FROM sellers
WHERE id = $1
RETURNING *;

-- name: RestoreSellerByID :one
UPDATE sellers
SET is_active  = TRUE,
    updated_at = NOW()
WHERE id = $1
  AND is_active = FALSE
RETURNING *;

-- name: UpdateSellerByID :one
UPDATE sellers
SET name        = COALESCE(sqlc.narg('name'), name),
    pix         = COALESCE(sqlc.narg('pix'), pix),
    email       = COALESCE(sqlc.narg('email'), email),
    phone       = COALESCE(sqlc.narg('phone'), phone),
    observation = COALESCE(sqlc.narg('observation'), observation),
    cpf         = COALESCE(sqlc.narg('cpf'), cpf),
    updated_at  = NOW()
WHERE id = $1
RETURNING *;
