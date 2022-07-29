package storage

import (
	"github.com/coreyvan/go-address/internal/service"
	expand "github.com/openvenues/gopostal/expand"
	"log"
)

type postgresStorage struct{}

func NewPostgresStorage() *postgresStorage {
	return &postgresStorage{}
}

func (p *postgresStorage) GetAddressByID(id string) (service.Address, error) {
	return service.Address{
		ID:           id,
		StreetNumber: "1234",
		StreetName:   "Sesame St",
		City:         "New York",
		State:        "NY",
		Zipcode:      12345,
	}, nil
}

func (p *postgresStorage) GetAddressBySearch(query string) ([]service.AddressSearch, error) {
	return []service.AddressSearch{
		{
			Address: service.Address{
				ID:           "1234",
				StreetNumber: "1234",
				StreetName:   "Sesame St",
				City:         "New York",
				State:        "NY",
				Zipcode:      12345,
			},
			Similarity: 1,
		},
	}, nil
}

func (p *postgresStorage) CreateAddress(address service.CreateAddress) (service.Address, error) {
	// TODO: actually create this in pg
	addr := service.Address{
		ID:           "1234",
		StreetNumber: address.StreetNumber,
		StreetName:   address.StreetName,
		City:         address.City,
		State:        address.State,
		Zipcode:      address.Zipcode,
		Geo:          address.Geo,
	}

	if err := p.createAddressSearchTerms(addr); err != nil {
		return service.Address{}, err
	}

	return addr, nil
}

func (p *postgresStorage) createAddressSearchTerms(address service.Address) error {
	addrStr := address.FormattedString()

	terms := expand.ExpandAddress(addrStr)

	return p.bulkAddTerms(terms, address.ID)
}

func (p *postgresStorage) bulkAddTerms(terms []string, id string) error {
	// TODO: bulk add terms to pg
	for _, t := range terms {
		log.Printf("mapping %s to id %s", t, id)
	}

	return nil
}
