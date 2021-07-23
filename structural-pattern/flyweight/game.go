package flyweight

type Game struct {
	Terrorists        []*Player
	CounterTerrorists []*Player
}

func (c *Game) AddTerrorist(dressType string) {
	player := NewPlayer("T", dressType)
	c.Terrorists = append(c.Terrorists, player)
}

func (c *Game) AddCounterTerrorist(dressType string) {
	player := NewPlayer("CT", dressType)
	c.CounterTerrorists = append(c.CounterTerrorists, player)
}
