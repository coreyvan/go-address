package service

import "context"

type Service interface {
	GetAddressByID(ctx context.Context, id int32) (Address, error)
	GetAddressByString(ctx context.Context, address string) (Address, error)
	GetAddressBySearch(ctx context.Context, query string) ([]AddressSearch, error)
	CreateAddress(ctx context.Context, address CreateAddress) (Address, error)
}

type service struct {
	storage Storage
}

type Storage interface {
	GetAddressByID(ctx context.Context, id int32) (Address, error)
	GetAddressByString(ctx context.Context, address string) (Address, error)
	GetAddressBySearch(ctx context.Context, query string) ([]AddressSearch, error)
	CreateAddress(ctx context.Context, address CreateAddress) (Address, error)
}

func NewService(storage Storage) *service {
	return &service{
		storage: storage,
	}
}

func (s *service) GetAddressByID(ctx context.Context, id int32) (Address, error) {
	// TODO validation logic
	return s.storage.GetAddressByID(ctx, id)
}

func (s *service) GetAddressByString(ctx context.Context, address string) (Address, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetAddressBySearch(ctx context.Context, query string) ([]AddressSearch, error) {
	return s.storage.GetAddressBySearch(ctx, query)
}

func (s *service) CreateAddress(ctx context.Context, address CreateAddress) (Address, error) {
	// TODO validation logic
	return s.storage.CreateAddress(ctx, address)
}
