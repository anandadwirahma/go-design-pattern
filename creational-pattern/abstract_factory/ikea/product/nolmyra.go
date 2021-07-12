package product

type NOLMYRA struct {
}

func (n NOLMYRA) Price() float64 {
	return 595000
}

func (n NOLMYRA) IsIotEnabled() bool {
	return false
}

func (n NOLMYRA) IsSoft() bool {
	return false
}
