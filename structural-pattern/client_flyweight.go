package main

import (
	"fmt"
	"go-design-pattern/structural-pattern/flyweight"
)

func CallFlyweightMethod() {
	game := flyweight.Game{}

	//Add Terrorist
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(flyweight.CounterTerroristDressType)
	game.AddCounterTerrorist(flyweight.CounterTerroristDressType)
	game.AddCounterTerrorist(flyweight.CounterTerroristDressType)
	game.AddCounterTerrorist(flyweight.CounterTerroristDressType)
	game.AddCounterTerrorist(flyweight.CounterTerroristDressType)

	dressFactoryInstance := flyweight.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
