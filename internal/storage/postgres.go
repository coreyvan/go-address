package storage

type postgresStorage struct{}

func NewPostgresStorage() *postgresStorage {
	return &postgresStorage{}
}
