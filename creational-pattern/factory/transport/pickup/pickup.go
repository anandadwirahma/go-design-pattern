package pickup

type Pickup struct{}

func (p Pickup) GetTransport() string {
	return "Pickup"
}
