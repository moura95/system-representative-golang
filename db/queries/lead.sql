-- name: CreateLead :one
INSERT INTO leads (name, email, phone, origin)
VALUES ($1, $2, $3, $4)
    RETURNING *;

-- name: GetLeadByID :one
SELECT *
FROM leads
WHERE id = $1;

-- name: GetLeadByEmail :one
SELECT *
FROM leads
WHERE email = $1;


-- name: ListLeads :many
SELECT *
FROM leads;

-- name: DeleteLeadByID :one
DELETE
FROM leads
WHERE id = $1
    RETURNING *;