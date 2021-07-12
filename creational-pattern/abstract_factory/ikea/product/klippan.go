package product

import "go-design-pattern/creational-pattern/abstract_factory"

type KLIPPAN struct {
}

func (k KLIPPAN) Price() float64 {
	return 2895000
}
func (k KLIPPAN) Style() abstract_factory.SofaStyle {
	return abstract_factory.EuropeanStyle
}
