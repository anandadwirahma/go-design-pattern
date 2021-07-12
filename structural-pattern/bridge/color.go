package bridge

type Color interface {
	Hex() string
}

type Red struct{}

func (r Red) Hex() string {
	return "#Red"
}

type Green struct{}

func (g Green) Hex() string {
	return "#Green"
}
