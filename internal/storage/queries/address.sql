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

-- name: SearchAddresses :many
SELECT a.id, a.street_number, a.street_name, a.city, a.state, a.zipcode, max(similarity(address, sqlc.arg(query))) as sim from address_lookups
JOIN addresses a ON address_lookups.address_id = a.id
GROUP BY a.id, a.street_number, a.street_name, a.city, a.state, a.zipcode
ORDER BY sim DESC;