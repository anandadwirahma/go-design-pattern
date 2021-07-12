package product

type BeanBag struct {
	PriceItem     float64
	SoftnessLevel int
}

func (b BeanBag) Price() float64 {
	return b.PriceItem
}

func (b BeanBag) IsIotEnabled() bool {
	return false
}

func (b BeanBag) IsSoft() bool {
	return b.SoftnessLevel > 10
}
