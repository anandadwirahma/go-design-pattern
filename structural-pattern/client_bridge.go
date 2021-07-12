package main

import (
	"fmt"
	"go-design-pattern/structural-pattern/bridge"
)

func CallBridgeMethod() {
	circle := bridge.Circle{
		Radius: 10,
	}
	circle.Draw()
	fmt.Printf("Area: %f\n", circle.Area())

	greenCircle := bridge.Circle{
		Color:  bridge.Green{},
		Radius: 10,
	}
	greenCircle.Draw()
	fmt.Printf("Area: %f\n", greenCircle.Area())

	fmt.Println()

	square := bridge.Square{
		Length: 10,
	}
	square.Draw()
	fmt.Printf("Area: %f\n", square.Area())

	redSquare := square.Clone() //Combine with Prototype pattern
	redSquare.Color = bridge.Red{}
	redSquare.Draw()
	fmt.Printf("Area: %f\n", redSquare.Area())
}
