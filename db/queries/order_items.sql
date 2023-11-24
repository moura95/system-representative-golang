-- name: CreateOrderItems :one
INSERT INTO order_items (order_id, product_id, quantity, price, discount)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteOrderItemsByID :exec
DELETE
FROM order_items
WHERE order_id = $1
  AND product_id = $2;

-- name: GetOrderItemsByID :one
SELECT p.name        AS product_name,
       p.ipi         AS ipi,
       p.description AS description,
       p.code        AS code,
       oi.*
FROM order_items oi
         JOIN products p on oi.product_id = p.id
WHERE oi.order_id = $1
  AND oi.product_id = $2;

-- name: ListOrdersItemsByOrderID :many
SELECT p.name        AS product_name,
       p.ipi         AS ipi,
       p.description AS description,
       p.code        AS code,
       oi.*
FROM order_items oi
         JOIN products p ON oi.product_id = p.id
WHERE oi.order_id = $1
ORDER BY oi.ctid;

-- name: UpdateOrderItemByID :one
UPDATE order_items
SET quantity = COALESCE(sqlc.narg('quantity'), quantity),
    price    = COALESCE(sqlc.narg('price'), price),
    discount = COALESCE(sqlc.narg('discount'), discount)
WHERE order_id = $1
  AND product_id = $2
RETURNING *;
