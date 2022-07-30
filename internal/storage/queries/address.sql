-- name: GetAddressByID :one
SELECT * FROM addresses
WHERE id = $1;