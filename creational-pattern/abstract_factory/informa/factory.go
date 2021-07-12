package informa

import (
	"go-design-pattern/creational-pattern/abstract_factory"
	"go-design-pattern/creational-pattern/abstract_factory/informa/product"
)

type Informa struct {
}

func (i Informa) CreateChair() abstract_factory.Chair {
	return &product.BeanBag{
		PriceItem:     1500000,
		SoftnessLevel: 100,
	}
}

func (i Informa) CreateCoffeeTable() abstract_factory.CoffeeTable {
	return nil
}

func (i Informa) CreateSofa() abstract_factory.Sofa {
	return nil
}
