// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: product.sql

package repository

import (
	"context"
	"database/sql"
	"time"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (representative_id, factory_id, name, code, price, ipi, reference, description, image_url)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, representative_id, factory_id, name, code, price, ipi, reference, description, image_url, is_active, created_at, updated_at
`

type CreateProductParams struct {
	RepresentativeID int32
	FactoryID        int32
	Name             string
	Code             string
	Price            string
	Ipi              sql.NullString
	Reference        sql.NullString
	Description      sql.NullString
	ImageUrl         sql.NullString
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.RepresentativeID,
		arg.FactoryID,
		arg.Name,
		arg.Code,
		arg.Price,
		arg.Ipi,
		arg.Reference,
		arg.Description,
		arg.ImageUrl,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.RepresentativeID,
		&i.FactoryID,
		&i.Name,
		&i.Code,
		&i.Price,
		&i.Ipi,
		&i.Reference,
		&i.Description,
		&i.ImageUrl,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProductByID = `-- name: DeleteProductByID :one
UPDATE products
SET is_active  = FALSE,
    updated_at = NOW()
WHERE id = $1
  AND is_active = TRUE
RETURNING id, representative_id, factory_id, name, code, price, ipi, reference, description, image_url, is_active, created_at, updated_at
`

func (q *Queries) DeleteProductByID(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, deleteProductByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.RepresentativeID,
		&i.FactoryID,
		&i.Name,
		&i.Code,
		&i.Price,
		&i.Ipi,
		&i.Reference,
		&i.Description,
		&i.ImageUrl,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProductByID = `-- name: GetProductByID :one
SELECT products.id, products.representative_id, products.factory_id, products.name, products.code, products.price, products.ipi, products.reference, products.description, products.image_url, products.is_active, products.created_at, products.updated_at, companies.name AS factory_name
FROM products
         JOIN companies ON products.factory_id = companies.id
WHERE products.id = $1
`

type GetProductByIDRow struct {
	ID               int32
	RepresentativeID int32
	FactoryID        int32
	Name             string
	Code             string
	Price            string
	Ipi              sql.NullString
	Reference        sql.NullString
	Description      sql.NullString
	ImageUrl         sql.NullString
	IsActive         bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
	FactoryName      string
}

func (q *Queries) GetProductByID(ctx context.Context, id int32) (GetProductByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getProductByID, id)
	var i GetProductByIDRow
	err := row.Scan(
		&i.ID,
		&i.RepresentativeID,
		&i.FactoryID,
		&i.Name,
		&i.Code,
		&i.Price,
		&i.Ipi,
		&i.Reference,
		&i.Description,
		&i.ImageUrl,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FactoryName,
	)
	return i, err
}

const listProductsByRepresentativeID = `-- name: ListProductsByRepresentativeID :many
SELECT products.id, products.representative_id, products.factory_id, products.name, products.code, products.price, products.ipi, products.reference, products.description, products.image_url, products.is_active, products.created_at, products.updated_at, companies.name AS factory_name
FROM products
         JOIN companies ON products.factory_id = companies.id
WHERE products.representative_id = $1
  AND products.is_active = $2
  AND ($3::int IS NULL OR $3 = products.factory_id)
ORDER BY products.id DESC
`

type ListProductsByRepresentativeIDParams struct {
	RepresentativeID int32
	IsActive         bool
	FactoryID        sql.NullInt32
}

type ListProductsByRepresentativeIDRow struct {
	ID               int32
	RepresentativeID int32
	FactoryID        int32
	Name             string
	Code             string
	Price            string
	Ipi              sql.NullString
	Reference        sql.NullString
	Description      sql.NullString
	ImageUrl         sql.NullString
	IsActive         bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
	FactoryName      string
}

func (q *Queries) ListProductsByRepresentativeID(ctx context.Context, arg ListProductsByRepresentativeIDParams) ([]ListProductsByRepresentativeIDRow, error) {
	rows, err := q.db.QueryContext(ctx, listProductsByRepresentativeID, arg.RepresentativeID, arg.IsActive, arg.FactoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProductsByRepresentativeIDRow{}
	for rows.Next() {
		var i ListProductsByRepresentativeIDRow
		if err := rows.Scan(
			&i.ID,
			&i.RepresentativeID,
			&i.FactoryID,
			&i.Name,
			&i.Code,
			&i.Price,
			&i.Ipi,
			&i.Reference,
			&i.Description,
			&i.ImageUrl,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FactoryName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeProductByID = `-- name: RemoveProductByID :one
DELETE
FROM products
WHERE id = $1
RETURNING id, representative_id, factory_id, name, code, price, ipi, reference, description, image_url, is_active, created_at, updated_at
`

func (q *Queries) RemoveProductByID(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, removeProductByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.RepresentativeID,
		&i.FactoryID,
		&i.Name,
		&i.Code,
		&i.Price,
		&i.Ipi,
		&i.Reference,
		&i.Description,
		&i.ImageUrl,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const restoreProductByID = `-- name: RestoreProductByID :one
UPDATE products
SET is_active  = TRUE,
    updated_at = NOW()
WHERE id = $1
  AND is_active = FALSE
RETURNING id, representative_id, factory_id, name, code, price, ipi, reference, description, image_url, is_active, created_at, updated_at
`

func (q *Queries) RestoreProductByID(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, restoreProductByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.RepresentativeID,
		&i.FactoryID,
		&i.Name,
		&i.Code,
		&i.Price,
		&i.Ipi,
		&i.Reference,
		&i.Description,
		&i.ImageUrl,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProductByID = `-- name: UpdateProductByID :one
UPDATE products
SET name        = COALESCE($2, name),
    code        = COALESCE($3, code),
    price       = COALESCE($4, price),
    ipi         = COALESCE($5, ipi),
    reference   = COALESCE($6, reference),
    description = COALESCE($7, description),
    image_url   = COALESCE($8, image_url),
    updated_at  = NOW()
WHERE id = $1
RETURNING id, representative_id, factory_id, name, code, price, ipi, reference, description, image_url, is_active, created_at, updated_at
`

type UpdateProductByIDParams struct {
	ID          int32
	Name        sql.NullString
	Code        sql.NullString
	Price       sql.NullString
	Ipi         sql.NullString
	Reference   sql.NullString
	Description sql.NullString
	ImageUrl    sql.NullString
}

func (q *Queries) UpdateProductByID(ctx context.Context, arg UpdateProductByIDParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProductByID,
		arg.ID,
		arg.Name,
		arg.Code,
		arg.Price,
		arg.Ipi,
		arg.Reference,
		arg.Description,
		arg.ImageUrl,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.RepresentativeID,
		&i.FactoryID,
		&i.Name,
		&i.Code,
		&i.Price,
		&i.Ipi,
		&i.Reference,
		&i.Description,
		&i.ImageUrl,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
