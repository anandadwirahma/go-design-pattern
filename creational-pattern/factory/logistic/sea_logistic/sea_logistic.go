package sea_logistic

import (
	"go-design-pattern/creational-pattern/factory"
	"go-design-pattern/creational-pattern/factory/transport/ship"
)

type SeaLogistic struct{}

func (s SeaLogistic) CreateTransport() factory.Transport {
	return &ship.Ship{}
}
