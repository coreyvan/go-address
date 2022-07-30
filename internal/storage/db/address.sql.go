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
	row := q.db.QueryRowContext(ctx, createAddress,
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
	_, err := q.db.ExecContext(ctx, createAddressLookup, arg.Address, arg.AddressID)
	return err
}

const getAddressByID = `-- name: GetAddressByID :one
SELECT id, street_number, street_name, city, state, zipcode
FROM addresses
WHERE id = $1
`

func (q *Queries) GetAddressByID(ctx context.Context, id int32) (Address, error) {
	row := q.db.QueryRowContext(ctx, getAddressByID, id)
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
