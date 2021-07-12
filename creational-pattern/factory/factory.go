package factory

type Transport interface {
	GetTransport() string
}

type Logistic interface {
	CreateTransport() Transport
}
