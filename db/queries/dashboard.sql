-- name: TopBuyerByRepresentativeID :many
SELECT customer_name,
       cast(total_by_customer as decimal(10, 2))                                                                      as total_by_customer,
       cast((SELECT COALESCE(sum(total), 0)
             FROM orders
             WHERE status = 'Concluido'
               AND representative_id = sqlc.arg('representative_id')::int
               AND is_active = TRUE
               AND total > 0
               AND created_at BETWEEN sqlc.arg('start_date')::date AND sqlc.arg('end_date')::date) as decimal(10, 2)) as total_sales
FROM (SELECT companies.name AS customer_name, sum(orders.total) as total_by_customer
      FROM orders
               JOIN companies ON orders.customer_id = companies.id
      WHERE orders.status = 'Concluido'
        AND orders.representative_id = sqlc.arg('representative_id')::int
        AND orders.is_active = TRUE
        AND orders.total > 0
        AND orders.created_at BETWEEN sqlc.arg('start_date')::date AND sqlc.arg('end_date')::date
      GROUP BY orders.customer_id, companies.name
      ORDER BY total_by_customer DESC
      LIMIT sqlc.arg('top')::int) as top_5_customers;

-- name: TotalSalesPerDayByRepresentativeID :many
SELECT to_char(created_at, 'YYYY-MM-DD')                                                                              AS "day",
       CAST(COALESCE(sum(total), 0) AS decimal(10, 2))                                                                AS total_day,
       CAST((SELECT COALESCE(sum(total), 0)
             FROM orders
             WHERE status = 'Concluido'
               AND representative_id = sqlc.arg('representative_id')::int
               AND is_active = TRUE
               AND total > 0
               AND created_at BETWEEN sqlc.arg('start_date')::date AND sqlc.arg('end_date')::date) as decimal(10, 2)) as total
FROM orders
WHERE status = 'Concluido'
  AND representative_id = sqlc.arg('representative_id')::int
  AND is_active = TRUE
  AND total > 0
  AND created_at BETWEEN sqlc.arg('start_date')::date AND sqlc.arg('end_date')::date
GROUP BY "day";

-- name: TopSalesPerProductByRepresentativeID :many
SELECT p.name                                                                                                         as product_name,
       CAST(SUM(((oi.price * oi.quantity) * (1 - oi.discount / 100))) AS decimal(10, 2))                              AS sub_total,
       CAST((SELECT COALESCE(sum(total), 0)
             FROM orders
             WHERE status = 'Concluido'
               AND representative_id = sqlc.arg('representative_id')::int
               AND is_active = TRUE
               AND total > 0
               AND created_at BETWEEN sqlc.arg('start_date')::date AND sqlc.arg('end_date')::date) as decimal(10, 2)) as total
FROM order_items oi
         JOIN orders o ON oi.order_id = o.id
         JOIN products p ON oi.product_id = p.id
WHERE o.status = 'Concluido'
  AND o.representative_id = sqlc.arg('representative_id')::int
  AND o.is_active = TRUE
  AND o.total > 0
  AND o.created_at BETWEEN sqlc.arg('start_date')::date AND sqlc.arg('end_date')::date
GROUP BY p.id
ORDER BY sub_total DESC
LIMIT sqlc.arg('top')::int;

-- name: TopFactoryByRepresentativeID :many
SELECT f.name                                                                                                         as factory_name,
       CAST(SUM(((oi.price * oi.quantity) * (1 - oi.discount / 100))) AS decimal(10, 2))                              AS sub_total,
       CAST((SELECT COALESCE(sum(total), 10)
             FROM orders
             WHERE status = 'Concluido'
               AND representative_id = sqlc.arg('representative_id')::int
               AND is_active = TRUE
               AND total > 0
               AND created_at BETWEEN sqlc.arg('start_date')::date AND sqlc.arg('end_date')::date) as decimal(10, 2)) as total
FROM order_items oi
         JOIN orders o ON oi.order_id = o.id
         JOIN products p ON oi.product_id = p.id
         JOIN companies f ON p.factory_id = f.id
WHERE o.status = 'Concluido'
  AND o.representative_id = sqlc.arg('representative_id')::int
  AND o.is_active = TRUE
  AND o.total > 0
  AND o.created_at BETWEEN sqlc.arg('start_date')::date AND sqlc.arg('end_date')::date
GROUP BY f.id
ORDER BY sub_total DESC
LIMIT sqlc.arg('top')::int;
