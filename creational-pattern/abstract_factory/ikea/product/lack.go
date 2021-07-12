package product

import "go-design-pattern/creational-pattern/abstract_factory"

type LACK struct {
}

func (l LACK) Price() float64 {
	return 999000
}

func (l LACK) Size() abstract_factory.Dimension {
	return abstract_factory.Dimension{
		Length: 118,
		Width:  78,
		Height: 40,
	}
}

func (l LACK) IsFoldable() bool {
	return false
}
