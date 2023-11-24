-- name: CreatePaymentForm :one
INSERT INTO form_payments(name, representative_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetPaymentFormByID :one
SELECT *
FROM form_payments
WHERE id = $1;

-- name: ListPaymentFormsByRepresentativeID :many
SELECT *
FROM form_payments
WHERE representative_id = $1;

-- name: DeletePaymentFormByID :one
DELETE
FROM form_payments
WHERE id = $1
RETURNING *;

-- name: UpdatePaymentFormByID :one
UPDATE form_payments
SET name = $1
WHERE id = $2
RETURNING *;
