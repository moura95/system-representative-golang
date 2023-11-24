-- name: CreateCompany :one
INSERT INTO companies (representative_id, type, name, email, website, logo_url, street, number, city, state, zip_code,
                       cnpj,
                       fantasy_name, ie, phone)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
RETURNING *;

-- name: DeleteCompanyByID :one
UPDATE companies
SET is_active  = FALSE,
    updated_at = NOW()
WHERE id = $1
  AND is_active = TRUE
RETURNING *;

-- name: GetCompanyByID :one
SELECT *
FROM companies
WHERE id = $1;

-- name: GetCompanyUserByID :one
SELECT *
FROM companies c
         JOIN representatives r on r.id = c.representative_id
WHERE r.id = sqlc.arg('representativeID')
  AND c.id = sqlc.arg('companyID');

-- name: ListCompaniesByRepresentativeID :many
SELECT *
FROM companies
WHERE representative_id = $1
  AND type = $2
  AND is_active = $3;

-- name: RemoveCompanyByID :one
DELETE
FROM companies
WHERE id = $1
RETURNING *;

-- name: RestoreCompanyByID :one
UPDATE companies
SET is_active  = TRUE,
    updated_at = NOW()
WHERE id = $1
  AND is_active = FALSE
RETURNING *;

-- name: UpdateCompanyByID :one
UPDATE companies
SET cnpj        = COALESCE(sqlc.narg('cnpj'), cnpj),
    name        = COALESCE(sqlc.narg('name'), name),
    fantasy_name= COALESCE(sqlc.narg('fantasy_name'), fantasy_name),
    ie          = COALESCE(sqlc.narg('ie'), ie),
    phone       = COALESCE(sqlc.narg('phone'), phone),
    email       = COALESCE(sqlc.narg('email'), email),
    website     = COALESCE(sqlc.narg('website'), website),
    logo_url    = COALESCE(sqlc.narg('logo_url'), logo_url),
    zip_code    = COALESCE(sqlc.narg('zip_code'), zip_code),
    state       = COALESCE(sqlc.narg('state'), state),
    city        = COALESCE(sqlc.narg('city'), city),
    street      = COALESCE(sqlc.narg('street'), street),
    number      = COALESCE(sqlc.narg('number'), number),
    updated_at  = NOW()
WHERE id = $1
RETURNING *;
