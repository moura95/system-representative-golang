-- name: CreateProduct :one
INSERT INTO products (representative_id, factory_id, name, code, price, ipi, reference, description, image_url)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: DeleteProductByID :one
UPDATE products
SET is_active  = FALSE,
    updated_at = NOW()
WHERE id = $1
  AND is_active = TRUE
RETURNING *;

-- name: GetProductByID :one
SELECT products.*, companies.name AS factory_name
FROM products
         JOIN companies ON products.factory_id = companies.id
WHERE products.id = $1;

-- name: ListProductsByRepresentativeID :many
SELECT products.*, companies.name AS factory_name
FROM products
         JOIN companies ON products.factory_id = companies.id
WHERE products.representative_id = $1
  AND products.is_active = $2
  AND (sqlc.narg('factory_id')::int IS NULL OR sqlc.narg('factory_id') = products.factory_id)
ORDER BY products.id DESC;

-- name: RemoveProductByID :one
DELETE
FROM products
WHERE id = $1
RETURNING *;

-- name: RestoreProductByID :one
UPDATE products
SET is_active  = TRUE,
    updated_at = NOW()
WHERE id = $1
  AND is_active = FALSE
RETURNING *;

-- name: UpdateProductByID :one
UPDATE products
SET name        = COALESCE(sqlc.narg('name'), name),
    code        = COALESCE(sqlc.narg('code'), code),
    price       = COALESCE(sqlc.narg('price'), price),
    ipi         = COALESCE(sqlc.narg('ipi'), ipi),
    reference   = COALESCE(sqlc.narg('reference'), reference),
    description = COALESCE(sqlc.narg('description'), description),
    image_url   = COALESCE(sqlc.narg('image_url'), image_url),
    updated_at  = NOW()
WHERE id = $1
RETURNING *;
