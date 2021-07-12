package main

import (
	"fmt"

	"go-design-pattern/creational-pattern/builder"
	"go-design-pattern/creational-pattern/builder/ClassicHouse"
	"go-design-pattern/creational-pattern/builder/ModernHouse"
)

func CallBuilderMethod() {
	classicHouseBuilder := &ClassicHouse.ClassicHouse{}
	modernHouseBuilder := &ModernHouse.ModernHouse{}

	// Classic House Builder
	director := builder.NewDirector(classicHouseBuilder)
	house := director.BuildHouse()

	fmt.Println("*** Classic House ***")
	fmt.Printf("Floor Type: %s\n", house.FloorType)
	fmt.Printf("Window Type: %s\n", house.WindowType)
	fmt.Printf("Num Of Window: %d\n", house.NumOfWindows)
	fmt.Printf("Num Of Door: %d\n", house.NumOfDoors)
	fmt.Printf("Swimming Pool: %t\n", house.HasSwimmingPool)

	// Modern House Builder
	director = builder.NewDirector(modernHouseBuilder)
	house = director.BuildHouse()

	fmt.Println("\n*** Modern House ***")
	fmt.Printf("Floor Type: %s\n", house.FloorType)
	fmt.Printf("Window Type: %s\n", house.WindowType)
	fmt.Printf("Num Of Window: %d\n", house.NumOfWindows)
	fmt.Printf("Num Of Door: %d\n", house.NumOfDoors)
	fmt.Printf("Swimming Pool: %t\n", house.HasSwimmingPool)
}
