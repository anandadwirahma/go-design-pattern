package ikea

import (
	"go-design-pattern/creational-pattern/abstract_factory"
	"go-design-pattern/creational-pattern/abstract_factory/ikea/product"
)

type Ikea struct {
}

func (i Ikea) CreateChair() abstract_factory.Chair {
	return &product.NOLMYRA{}
}

func (i Ikea) CreateCoffeeTable() abstract_factory.CoffeeTable {
	return &product.LACK{}
}

func (i Ikea) CreateSofa() abstract_factory.Sofa {
	return &product.KLIPPAN{}
}
