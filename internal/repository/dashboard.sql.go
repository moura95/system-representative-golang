// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: dashboard.sql

package repository

import (
	"context"
	"time"
)

const topBuyerByRepresentativeID = `-- name: TopBuyerByRepresentativeID :many
SELECT customer_name,
       cast(total_by_customer as decimal(10, 2))                                                                      as total_by_customer,
       cast((SELECT COALESCE(sum(total), 0)
             FROM orders
             WHERE status = 'Concluido'
               AND representative_id = $1::int
               AND is_active = TRUE
               AND total > 0
               AND created_at BETWEEN $2::date AND $3::date) as decimal(10, 2)) as total_sales
FROM (SELECT companies.name AS customer_name, sum(orders.total) as total_by_customer
      FROM orders
               JOIN companies ON orders.customer_id = companies.id
      WHERE orders.status = 'Concluido'
        AND orders.representative_id = $1::int
        AND orders.is_active = TRUE
        AND orders.total > 0
        AND orders.created_at BETWEEN $2::date AND $3::date
      GROUP BY orders.customer_id, companies.name
      ORDER BY total_by_customer DESC
      LIMIT $4::int) as top_5_customers
`

type TopBuyerByRepresentativeIDParams struct {
	RepresentativeID int32
	StartDate        time.Time
	EndDate          time.Time
	Top              int32
}

type TopBuyerByRepresentativeIDRow struct {
	CustomerName    string
	TotalByCustomer string
	TotalSales      string
}

func (q *Queries) TopBuyerByRepresentativeID(ctx context.Context, arg TopBuyerByRepresentativeIDParams) ([]TopBuyerByRepresentativeIDRow, error) {
	rows, err := q.db.QueryContext(ctx, topBuyerByRepresentativeID,
		arg.RepresentativeID,
		arg.StartDate,
		arg.EndDate,
		arg.Top,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TopBuyerByRepresentativeIDRow{}
	for rows.Next() {
		var i TopBuyerByRepresentativeIDRow
		if err := rows.Scan(&i.CustomerName, &i.TotalByCustomer, &i.TotalSales); err != nil {
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

const topFactoryByRepresentativeID = `-- name: TopFactoryByRepresentativeID :many
SELECT f.name                                                                                                         as factory_name,
       CAST(SUM(((oi.price * oi.quantity) * (1 - oi.discount / 100))) AS decimal(10, 2))                              AS sub_total,
       CAST((SELECT COALESCE(sum(total), 10)
             FROM orders
             WHERE status = 'Concluido'
               AND representative_id = $1::int
               AND is_active = TRUE
               AND total > 0
               AND created_at BETWEEN $2::date AND $3::date) as decimal(10, 2)) as total
FROM order_items oi
         JOIN orders o ON oi.order_id = o.id
         JOIN products p ON oi.product_id = p.id
         JOIN companies f ON p.factory_id = f.id
WHERE o.status = 'Concluido'
  AND o.representative_id = $1::int
  AND o.is_active = TRUE
  AND o.total > 0
  AND o.created_at BETWEEN $2::date AND $3::date
GROUP BY f.id
ORDER BY sub_total DESC
LIMIT $4::int
`

type TopFactoryByRepresentativeIDParams struct {
	RepresentativeID int32
	StartDate        time.Time
	EndDate          time.Time
	Top              int32
}

type TopFactoryByRepresentativeIDRow struct {
	FactoryName string
	SubTotal    string
	Total       string
}

func (q *Queries) TopFactoryByRepresentativeID(ctx context.Context, arg TopFactoryByRepresentativeIDParams) ([]TopFactoryByRepresentativeIDRow, error) {
	rows, err := q.db.QueryContext(ctx, topFactoryByRepresentativeID,
		arg.RepresentativeID,
		arg.StartDate,
		arg.EndDate,
		arg.Top,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TopFactoryByRepresentativeIDRow{}
	for rows.Next() {
		var i TopFactoryByRepresentativeIDRow
		if err := rows.Scan(&i.FactoryName, &i.SubTotal, &i.Total); err != nil {
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

const topSalesPerProductByRepresentativeID = `-- name: TopSalesPerProductByRepresentativeID :many
SELECT p.name                                                                                                         as product_name,
       CAST(SUM(((oi.price * oi.quantity) * (1 - oi.discount / 100))) AS decimal(10, 2))                              AS sub_total,
       CAST((SELECT COALESCE(sum(total), 0)
             FROM orders
             WHERE status = 'Concluido'
               AND representative_id = $1::int
               AND is_active = TRUE
               AND total > 0
               AND created_at BETWEEN $2::date AND $3::date) as decimal(10, 2)) as total
FROM order_items oi
         JOIN orders o ON oi.order_id = o.id
         JOIN products p ON oi.product_id = p.id
WHERE o.status = 'Concluido'
  AND o.representative_id = $1::int
  AND o.is_active = TRUE
  AND o.total > 0
  AND o.created_at BETWEEN $2::date AND $3::date
GROUP BY p.id
ORDER BY sub_total DESC
LIMIT $4::int
`

type TopSalesPerProductByRepresentativeIDParams struct {
	RepresentativeID int32
	StartDate        time.Time
	EndDate          time.Time
	Top              int32
}

type TopSalesPerProductByRepresentativeIDRow struct {
	ProductName string
	SubTotal    string
	Total       string
}

func (q *Queries) TopSalesPerProductByRepresentativeID(ctx context.Context, arg TopSalesPerProductByRepresentativeIDParams) ([]TopSalesPerProductByRepresentativeIDRow, error) {
	rows, err := q.db.QueryContext(ctx, topSalesPerProductByRepresentativeID,
		arg.RepresentativeID,
		arg.StartDate,
		arg.EndDate,
		arg.Top,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TopSalesPerProductByRepresentativeIDRow{}
	for rows.Next() {
		var i TopSalesPerProductByRepresentativeIDRow
		if err := rows.Scan(&i.ProductName, &i.SubTotal, &i.Total); err != nil {
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

const totalSalesPerDayByRepresentativeID = `-- name: TotalSalesPerDayByRepresentativeID :many
SELECT to_char(created_at, 'YYYY-MM-DD')                                                                              AS "day",
       CAST(COALESCE(sum(total), 0) AS decimal(10, 2))                                                                AS total_day,
       CAST((SELECT COALESCE(sum(total), 0)
             FROM orders
             WHERE status = 'Concluido'
               AND representative_id = $1::int
               AND is_active = TRUE
               AND total > 0
               AND created_at BETWEEN $2::date AND $3::date) as decimal(10, 2)) as total
FROM orders
WHERE status = 'Concluido'
  AND representative_id = $1::int
  AND is_active = TRUE
  AND total > 0
  AND created_at BETWEEN $2::date AND $3::date
GROUP BY "day"
`

type TotalSalesPerDayByRepresentativeIDParams struct {
	RepresentativeID int32
	StartDate        time.Time
	EndDate          time.Time
}

type TotalSalesPerDayByRepresentativeIDRow struct {
	Day      string
	TotalDay string
	Total    string
}

func (q *Queries) TotalSalesPerDayByRepresentativeID(ctx context.Context, arg TotalSalesPerDayByRepresentativeIDParams) ([]TotalSalesPerDayByRepresentativeIDRow, error) {
	rows, err := q.db.QueryContext(ctx, totalSalesPerDayByRepresentativeID, arg.RepresentativeID, arg.StartDate, arg.EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TotalSalesPerDayByRepresentativeIDRow{}
	for rows.Next() {
		var i TotalSalesPerDayByRepresentativeIDRow
		if err := rows.Scan(&i.Day, &i.TotalDay, &i.Total); err != nil {
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
