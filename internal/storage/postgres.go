package storage

import (
	"context"
	"github.com/coreyvan/go-address/internal/service"
	"github.com/coreyvan/go-address/internal/storage/db"
	expand "github.com/openvenues/gopostal/expand"
	"github.com/pkg/errors"
)

type postgresStorage struct {
	db *db.Queries
}

func NewPostgresStorage() *postgresStorage {
	return &postgresStorage{}
}

func (p *postgresStorage) GetAddressByID(ctx context.Context, id int32) (service.Address, error) {
	addr, err := p.db.GetAddressByID(ctx, id)
	if err != nil {
		return service.Address{}, errors.Wrap(err, "getting address from pg")
	}

	return service.Address{
		ID:           addr.ID,
		StreetNumber: addr.StreetNumber,
		StreetName:   addr.StreetName,
		City:         addr.City,
		State:        addr.State,
		Zipcode:      addr.Zipcode,
	}, nil
}

func (p *postgresStorage) GetAddressBySearch(query string) ([]service.AddressSearch, error) {
	return []service.AddressSearch{
		{
			Address: service.Address{
				ID:           1234,
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

func (p *postgresStorage) CreateAddress(ctx context.Context, address service.CreateAddress) (service.Address, error) {
	addr, err := p.db.CreateAddress(ctx, db.CreateAddressParams{
		StreetNumber: address.StreetNumber,
		StreetName:   address.StreetName,
		City:         address.City,
		State:        address.State,
		Zipcode:      address.Zipcode,
	})
	if err != nil {
		return service.Address{}, errors.Wrap(err, "creating address in pg")
	}

	out := service.Address{
		ID:           addr.ID,
		StreetNumber: addr.StreetNumber,
		StreetName:   addr.StreetName,
		City:         addr.City,
		State:        addr.State,
		Zipcode:      addr.Zipcode,
	}

	if err := p.createAddressSearchTerms(ctx, out); err != nil {
		return service.Address{}, err
	}

	return out, nil
}

func (p *postgresStorage) createAddressSearchTerms(ctx context.Context, address service.Address) error {
	addrStr := address.FormattedString()

	terms := expand.ExpandAddress(addrStr)

	return p.bulkAddTerms(ctx, terms, address.ID)
}

func (p *postgresStorage) bulkAddTerms(ctx context.Context, terms []string, id int32) error {
	// TODO: bulk add terms to pg
	for _, t := range terms {
		if err := p.db.CreateAddressLookup(ctx, db.CreateAddressLookupParams{
			Address:   t,
			AddressID: id,
		}); err != nil {
			return errors.Wrap(err, "creating address lookup in pg")
		}
	}

	return nil
}
