// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import ()

type Address struct {
	ID           int32
	StreetNumber string
	StreetName   string
	City         string
	State        string
	Zipcode      int32
}

type AddressLookup struct {
	ID        int32
	Address   string
	AddressID int32
}
