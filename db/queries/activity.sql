-- name: CreateActivity :one
INSERT INTO activity (action, reference_url, user_id, representative_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListActivityByUserID :many
SELECT *
FROM activity
WHERE user_id = $1;

-- name: ListActivityByRepresentativeID :many
SELECT *
FROM activity
WHERE representative_id = $1;

-- name: ListActivity :many
SELECT *
FROM activity;