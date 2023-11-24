-- name: CreatePaymentReceipt :one
INSERT INTO payment_receipt (representative_id,status, type_payment,description, amount, expiration_date,payment_date,doc_number,recipient,payment_form,installment,interval_days,additional_info)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11,$12,$13)
    RETURNING *;

-- name: GetPaymentReceiptByID :one
SELECT pr.*,
       json_agg(json_build_object('file_id', fpr.id, 'payment_receipt_id', fpr.id,  'url_file', fpr.url_file)) AS files
FROM payment_receipt AS pr
         LEFT JOIN files_payment_receipt AS fpr ON pr.id = fpr.payment_receipt_id
WHERE pr.id = $1 and pr.representative_id = $2
GROUP BY pr.id;


-- name: ListPaymentReceiptByRepresentativeID :many
SELECT pr.*,
       json_agg(json_build_object('file_id', fpr.id, 'payment_receipt_id', fpr.id,  'url_file', fpr.url_file)) AS files
FROM payment_receipt AS pr
         LEFT JOIN files_payment_receipt AS fpr ON pr.id = fpr.payment_receipt_id
WHERE pr.representative_id = $1
GROUP BY pr.id;

-- name: DeletePaymentReceiptByID :one
DELETE
FROM payment_receipt
WHERE id = $1
    RETURNING *;

-- name: UpdatePaymentPaymentReceiptByID :one
UPDATE payment_receipt
SET description = COALESCE(sqlc.narg('description'),description) ,
amount = COALESCE(sqlc.narg('amount'),amount),
expiration_date = COALESCE(sqlc.narg('expiration_date'),expiration_date),
payment_date = COALESCE(sqlc.narg('payment_date'),payment_date),
doc_number = COALESCE(sqlc.narg('doc_number'),doc_number),
recipient = COALESCE(sqlc.narg('recipient'),recipient),
payment_form = COALESCE(sqlc.narg('payment_form'),payment_form),
status = COALESCE(sqlc.narg('status'),status),
additional_info = COALESCE(sqlc.narg('additional_info'),additional_info)
WHERE id = $1
    RETURNING *;


-- name: UploadFilePaymentReceipt :one
INSERT INTO files_payment_receipt (payment_receipt_id, url_file)
VALUES ($1, $2)
    RETURNING *;

-- name: DeleteFilePaymentReceiptByID :one
DELETE
FROM files_payment_receipt
WHERE id = $1
    RETURNING *;