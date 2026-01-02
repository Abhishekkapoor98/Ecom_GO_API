-- name: ListProducts :many
SELECT
    *
FROM
    products;

-- name: FindProductsByID :one
SELECT
    *
FROM 
    products
WHERE
    id = $1;