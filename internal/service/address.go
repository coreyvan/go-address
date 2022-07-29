package service

import "fmt"

type Address struct {
	ID           string `json:"id,omitempty"`
	StreetNumber string `json:"streetNumber,omitempty"`
	StreetName   string `json:"streetName,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zipcode      int    `json:"zipcode,omitempty"`
	Geo          *Geo   `json:"geo,omitempty"`
}

type Geo struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func (a Address) FormattedString() string {
	return fmt.Sprintf("%s %s, %s, %s %d", a.StreetNumber, a.StreetName, a.City, a.State, a.Zipcode)
}

type CreateAddress struct {
	StreetNumber string `json:"streetNumber,omitempty"`
	StreetName   string `json:"streetName,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zipcode      int    `json:"zipcode,omitempty"`
	Geo          *Geo   `json:"geo,omitempty"`
}

type AddressSearch struct {
	Address    Address
	Similarity int
}
