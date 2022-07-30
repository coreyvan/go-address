// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: address.sql

package db

import (
	"context"
)

const createAddress = `-- name: CreateAddress :one
INSERT INTO addresses (street_number, street_name, city, state, zipcode)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, street_number, street_name, city, state, zipcode
`

type CreateAddressParams struct {
	StreetNumber string
	StreetName   string
	City         string
	State        string
	Zipcode      int32
}

func (q *Queries) CreateAddress(ctx context.Context, arg CreateAddressParams) (Address, error) {
	row := q.db.QueryRow(ctx, createAddress,
		arg.StreetNumber,
		arg.StreetName,
		arg.City,
		arg.State,
		arg.Zipcode,
	)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.StreetNumber,
		&i.StreetName,
		&i.City,
		&i.State,
		&i.Zipcode,
	)
	return i, err
}

const createAddressLookup = `-- name: CreateAddressLookup :exec
INSERT INTO address_lookups (address, address_id)
VALUES ($1, $2)
`

type CreateAddressLookupParams struct {
	Address   string
	AddressID int32
}

func (q *Queries) CreateAddressLookup(ctx context.Context, arg CreateAddressLookupParams) error {
	_, err := q.db.Exec(ctx, createAddressLookup, arg.Address, arg.AddressID)
	return err
}

const getAddressByID = `-- name: GetAddressByID :one
SELECT id, street_number, street_name, city, state, zipcode
FROM addresses
WHERE id = $1
`

func (q *Queries) GetAddressByID(ctx context.Context, id int32) (Address, error) {
	row := q.db.QueryRow(ctx, getAddressByID, id)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.StreetNumber,
		&i.StreetName,
		&i.City,
		&i.State,
		&i.Zipcode,
	)
	return i, err
}

const searchAddresses = `-- name: SearchAddresses :many
SELECT a.id, a.street_number, a.street_name, a.city, a.state, a.zipcode, max(similarity(address, $1)) as sim from address_lookups
JOIN addresses a ON address_lookups.address_id = a.id
GROUP BY a.id, a.street_number, a.street_name, a.city, a.state, a.zipcode
ORDER BY sim DESC
`

type SearchAddressesRow struct {
	ID           int32
	StreetNumber string
	StreetName   string
	City         string
	State        string
	Zipcode      int32
	Sim          interface{}
}

func (q *Queries) SearchAddresses(ctx context.Context, query string) ([]SearchAddressesRow, error) {
	rows, err := q.db.Query(ctx, searchAddresses, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchAddressesRow
	for rows.Next() {
		var i SearchAddressesRow
		if err := rows.Scan(
			&i.ID,
			&i.StreetNumber,
			&i.StreetName,
			&i.City,
			&i.State,
			&i.Zipcode,
			&i.Sim,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
