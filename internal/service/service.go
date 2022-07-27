package service

type Service interface{}

type service struct {
	storage Storage
}

type Storage interface{}

func NewService(storage Storage) *service {
	return &service{
		storage: storage,
	}
}
