-- name: CreateRepresentative :one
INSERT INTO representatives (cnpj, name, email, website, logo_url, street, number, city,
                             state, zip_code, fantasy_name, ie, phone)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING *;

-- name: DeleteRepresentativeByID :one
UPDATE representatives
SET is_active  = FALSE,
    updated_at = NOW()
WHERE id = $1
  AND is_active = TRUE
RETURNING *;

-- name: GetAllRepresentativesByID :many
SELECT *
FROM representatives;

-- name: GetRepresentativesByID :one
SELECT *
FROM representatives
WHERE id = $1;

-- name: GetPlanByRepresentativesID :one
SELECT plan
FROM representatives
WHERE id = $1;

-- name: GetTotalUsersByRepresentativesID :one
SELECT COUNT(*)
FROM users
WHERE representative_id = $1;

-- name: GetRepresentativeDateExpByID :one
SELECT data_expire
FROM representatives
WHERE id = $1;

-- name: RemoveRepresentativeByID :one
DELETE
FROM representatives
WHERE id = $1
RETURNING *;

-- name: RestoreRepresentativeByID :one
UPDATE representatives
SET is_active  = TRUE,
    updated_at = NOW()
WHERE id = $1
  AND is_active = FALSE
RETURNING *;

-- name: UpdateRepresentativeByID :one
UPDATE representatives
SET name         = COALESCE(sqlc.narg('name'), name),
    email        = COALESCE(sqlc.narg('email'), email),
    website      = COALESCE(sqlc.narg('website'), website),
    logo_url     = COALESCE(sqlc.narg('logo_url'), logo_url),
    street       = COALESCE(sqlc.narg('street'), street),
    number       = COALESCE(sqlc.narg('number'), number),
    city         = COALESCE(sqlc.narg('city'), city),
    state        = COALESCE(sqlc.narg('state'), state),
    zip_code     = COALESCE(sqlc.narg('zip_code'), zip_code),
    cnpj         = COALESCE(sqlc.narg('cnpj'), cnpj),
    fantasy_name = COALESCE(sqlc.narg('fantasy_name'), fantasy_name),
    ie           = COALESCE(sqlc.narg('ie'), ie),
    phone        = COALESCE(sqlc.narg('phone'), phone),
    updated_at   = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdatePlanByID :one
UPDATE representatives
SET plan        = $2,
    is_active   = TRUE,
    stripe_id   = $3,
    data_expire = $4,
    updated_at  = NOW()
WHERE id = $1
RETURNING *;
