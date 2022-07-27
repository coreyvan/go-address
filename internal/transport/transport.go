package transport

type Transport interface {
	Listen() error
}
