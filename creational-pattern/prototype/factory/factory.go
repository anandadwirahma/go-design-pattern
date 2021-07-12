package factory

import (
	"fmt"

	"go-design-pattern/creational-pattern/prototype"
)

type Plane struct {
	Name       string
	Engine     string
	IsUseRadar bool
}

func (p Plane) Print() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Engine: %s\n", p.Engine)
	fmt.Printf("Is User Radar: %t\n", p.IsUseRadar)
}

func (p Plane) Clone() prototype.Plane {
	p.Name = p.Name + "_clone"
	p.Engine = p.Engine + "_clone"

	return p
}
