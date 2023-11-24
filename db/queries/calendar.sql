-- name: CreateCalendar :one
INSERT INTO calendars (title, visit_start, visit_end, allday, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteCalendarByID :one
DELETE
FROM calendars
WHERE id = $1
RETURNING *;

-- name: GetCalendarByID :one
SELECT *
FROM calendars
WHERE id = $1;

-- name: ListCalendarsByUserID :many
SELECT *
FROM calendars
WHERE user_id = $1;

-- name: UpdateCalendarByID :one
UPDATE calendars
SET title       = COALESCE(sqlc.narg('title'), title),
    visit_start = COALESCE(sqlc.narg('visit_start'), visit_start),
    visit_end   = COALESCE(sqlc.narg('visit_end'), visit_end),
    allday = COALESCE(sqlc.narg('allday'), allday)
WHERE id = $1
RETURNING *;
