-- name: CreateOrder :one
INSERT INTO orders (representative_id, factory_id, customer_id, portage_id, seller_id, form_payment_id,
                    order_number, url_pdf, buyer, shipping, status, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: DeleteOrderByID :one
UPDATE orders
SET is_active  = FALSE,
    updated_at = NOW()
WHERE id = $1
  AND is_active = TRUE
RETURNING *;

-- name: GetOrderByID :one
SELECT f.name  AS factory_name,
       c.name  AS customer_name,
       p.name  AS portage_name,
       s.name  AS seller_name,
       s.email  AS seller_email,
       c.email  AS customer_email,
       fp.name AS form_payment_name,
       c.cnpj AS customer_cnpj,
       f.cnpj AS factory_cnpj,
       o.*
FROM orders o
         JOIN companies f
              ON o.factory_id = f.id
         JOIN companies c
              ON o.customer_id = c.id
         JOIN companies p
              ON o.portage_id = p.id
         JOIN sellers s
              ON o.seller_id = s.id
         LEFT JOIN form_payments fp
              ON o.form_payment_id = fp.id
WHERE o.id = $1;

-- name: GetLastOrderByRepresentativeID :one
SELECT order_number
from orders
WHERE representative_id = $1
ORDER BY id DESC
LIMIT 1;

-- name: ListOrdersByRepresentativeID :many
SELECT f.name  AS factory_name,
       c.name  AS customer_name,
       p.name  AS portage_name,
       s.name  AS seller_name,
       s.email  AS seller_email,
       c.email  AS customer_email,
       fp.name AS form_payment_name,
       o.*
FROM orders o
         JOIN companies f
              ON o.factory_id = f.id
         JOIN companies c
              ON o.customer_id = c.id
         JOIN companies p
              ON o.portage_id = p.id
         JOIN sellers s
              ON o.seller_id = s.id
         LEFT JOIN form_payments fp
              ON o.form_payment_id = fp.id
WHERE o.representative_id = $1
  AND o.is_active = $2
ORDER BY o.id DESC;

-- name: RemoveOrderByID :one
DELETE
FROM orders
WHERE id = $1
RETURNING *;

-- name: RestoreOrderByID :one
UPDATE orders
SET is_active  = TRUE,
    updated_at = NOW()
WHERE id = $1
  AND is_active = FALSE
RETURNING *;

-- name: UpdateOrderByID :one
UPDATE orders
SET factory_id      = COALESCE(sqlc.narg('factory_id'), factory_id),
    customer_id     = COALESCE(sqlc.narg('customer_id'), customer_id),
    portage_id      = COALESCE(sqlc.narg('portage_id'), portage_id),
    seller_id       = COALESCE(sqlc.narg('seller_id'), seller_id),
    form_payment_id = COALESCE(sqlc.narg('form_payment_id'), form_payment_id),
    url_pdf         = COALESCE(sqlc.narg('url_pdf'), url_pdf),
    buyer           = COALESCE(sqlc.narg('buyer'), buyer),
    shipping        = COALESCE(sqlc.narg('shipping'), shipping),
    status          = COALESCE(sqlc.narg('status'), status),
    expired_at      = COALESCE(sqlc.narg('expired_at'), expired_at),
    created_at      = COALESCE(sqlc.narg('created_at'), created_at),
    updated_at      = NOW()
WHERE id = $1
RETURNING *;
