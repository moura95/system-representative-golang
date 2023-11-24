// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: smtp.sql

package repository

import (
	"context"
	"database/sql"
)

const createSmtp = `-- name: CreateSmtp :one
INSERT INTO smtp (representative_id, email, password, server, port)
VALUES ($1, $2, $3, $4, $5)
RETURNING representative_id, is_active, email, password, server, port, created_at, updated_at
`

type CreateSmtpParams struct {
	RepresentativeID int32
	Email            string
	Password         string
	Server           string
	Port             int32
}

func (q *Queries) CreateSmtp(ctx context.Context, arg CreateSmtpParams) (Smtp, error) {
	row := q.db.QueryRowContext(ctx, createSmtp,
		arg.RepresentativeID,
		arg.Email,
		arg.Password,
		arg.Server,
		arg.Port,
	)
	var i Smtp
	err := row.Scan(
		&i.RepresentativeID,
		&i.IsActive,
		&i.Email,
		&i.Password,
		&i.Server,
		&i.Port,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteSmtpByRepresentativeID = `-- name: DeleteSmtpByRepresentativeID :one
DELETE
FROM smtp
WHERE representative_id = $1
RETURNING representative_id, is_active, email, password, server, port, created_at, updated_at
`

func (q *Queries) DeleteSmtpByRepresentativeID(ctx context.Context, representativeID int32) (Smtp, error) {
	row := q.db.QueryRowContext(ctx, deleteSmtpByRepresentativeID, representativeID)
	var i Smtp
	err := row.Scan(
		&i.RepresentativeID,
		&i.IsActive,
		&i.Email,
		&i.Password,
		&i.Server,
		&i.Port,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSmtpByRepresentativeID = `-- name: GetSmtpByRepresentativeID :one
SELECT representative_id, is_active, email, password, server, port, created_at, updated_at
FROM smtp
WHERE representative_id = $1
`

func (q *Queries) GetSmtpByRepresentativeID(ctx context.Context, representativeID int32) (Smtp, error) {
	row := q.db.QueryRowContext(ctx, getSmtpByRepresentativeID, representativeID)
	var i Smtp
	err := row.Scan(
		&i.RepresentativeID,
		&i.IsActive,
		&i.Email,
		&i.Password,
		&i.Server,
		&i.Port,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateSmtpByID = `-- name: UpdateSmtpByID :one
UPDATE smtp
SET email        = COALESCE($2, email),
    password     = COALESCE($3, password),
    server       = COALESCE($4, server),
    port         = COALESCE($5, port),
    updated_at   = NOW()
WHERE representative_id = $1
RETURNING representative_id, is_active, email, password, server, port, created_at, updated_at
`

type UpdateSmtpByIDParams struct {
	RepresentativeID int32
	Email            sql.NullString
	Password         sql.NullString
	Server           sql.NullString
	Port             sql.NullInt32
}

func (q *Queries) UpdateSmtpByID(ctx context.Context, arg UpdateSmtpByIDParams) (Smtp, error) {
	row := q.db.QueryRowContext(ctx, updateSmtpByID,
		arg.RepresentativeID,
		arg.Email,
		arg.Password,
		arg.Server,
		arg.Port,
	)
	var i Smtp
	err := row.Scan(
		&i.RepresentativeID,
		&i.IsActive,
		&i.Email,
		&i.Password,
		&i.Server,
		&i.Port,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
