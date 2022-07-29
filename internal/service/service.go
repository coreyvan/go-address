package service

type Service interface {
	GetAddressByID(id string) (Address, error)
	GetAddressByString(address string) (Address, error)
	GetAddressBySearch(query string) ([]AddressSearch, error)
	CreateAddress(address CreateAddress) (Address, error)
}

type service struct {
	storage Storage
}

type Storage interface {
	GetAddressByID(id string) (Address, error)
	GetAddressByString(address string) (Address, error)
	GetAddressBySearch(query string) ([]AddressSearch, error)
	CreateAddress(address CreateAddress) (Address, error)
}

func NewService(storage Storage) *service {
	return &service{
		storage: storage,
	}
}

func (s *service) GetAddressByID(id string) (Address, error) {
	// TODO validation logic
	return s.storage.GetAddressByID(id)
}

func (s *service) GetAddressByString(address string) (Address, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetAddressBySearch(query string) ([]AddressSearch, error) {
	return s.storage.GetAddressBySearch(query)
}

func (s *service) CreateAddress(address CreateAddress) (Address, error) {
	// TODO validation logic
	return s.storage.CreateAddress(address)
}
