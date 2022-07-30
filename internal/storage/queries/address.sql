-- name: GetAddressByID :one
SELECT *
FROM addresses
WHERE id = $1;

-- name: CreateAddress :one
INSERT INTO addresses (street_number, street_name, city, state, zipcode)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: CreateAddressLookup :exec
INSERT INTO address_lookups (address, address_id)
VALUES ($1, $2);