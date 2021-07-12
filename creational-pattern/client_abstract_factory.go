package main

import (
	"fmt"

	"go-design-pattern/creational-pattern/abstract_factory"
	"go-design-pattern/creational-pattern/abstract_factory/ikea"
	"go-design-pattern/creational-pattern/abstract_factory/informa"
)

func CallAbstractFactoryMethod() {
	var furniture abstract_factory.FurnitureFactory

	fmt.Println("-- IKEA Product --")
	furniture = ikea.Ikea{}
	if chair := furniture.CreateChair(); chair != nil {
		fmt.Println("**Chair IKEA")
		fmt.Printf("Price: %f\n", chair.Price())
		fmt.Printf("Is Soft: %t\n", chair.IsSoft())
		fmt.Printf("Is IOT: %t\n", chair.IsIotEnabled())
	} else {
		fmt.Println("Ikea hasn't implement chair")
	}

	if sofa := furniture.CreateSofa(); sofa != nil {
		fmt.Println("**Sofa IKEA")
		fmt.Printf("Price: %f\n", sofa.Price())
		fmt.Printf("Style: %s\n", sofa.Style())
	} else {
		fmt.Println("Ikea hasn't implement chair")
	}

	if coffeeTable := furniture.CreateCoffeeTable(); coffeeTable != nil {
		fmt.Println("**Coffee Table IKEA")
		fmt.Printf("Price: %f\n", coffeeTable.Price())
		fmt.Printf("Width: %d\n", coffeeTable.Size().Width)
		fmt.Printf("Length: %d\n", coffeeTable.Size().Length)
		fmt.Printf("Height: %d\n", coffeeTable.Size().Height)
		fmt.Printf("Is Flodable: %t\n", coffeeTable.IsFoldable())
	} else {
		fmt.Println("Ikea hasn't implement coffee table")
	}

	fmt.Println("\n-- INFORMA Product --")
	furniture = informa.Informa{}
	if chair := furniture.CreateChair(); chair != nil {
		fmt.Println("**Chair INFORMA")
		fmt.Printf("Price: %f\n", chair.Price())
		fmt.Printf("Is Soft: %t\n", chair.IsSoft())
		fmt.Printf("Is IOT: %t\n", chair.IsIotEnabled())
	} else {
		fmt.Println("INFORMA hasn't implement chair")
	}

	if sofa := furniture.CreateSofa(); sofa != nil {
		fmt.Println("**Sofa INFORMA")
		fmt.Printf("Price: %f\n", sofa.Price())
		fmt.Printf("Style: %s\n", sofa.Style())
	} else {
		fmt.Println("INFORMA hasn't implement chair")
	}

	if coffeeTable := furniture.CreateCoffeeTable(); coffeeTable != nil {
		fmt.Println("**Coffee Table INFORMA")
		fmt.Printf("Price: %f\n", coffeeTable.Price())
		fmt.Printf("Width: %d\n", coffeeTable.Size().Width)
		fmt.Printf("Length: %d\n", coffeeTable.Size().Length)
		fmt.Printf("Height: %d\n", coffeeTable.Size().Height)
		fmt.Printf("Is Flodable: %t\n", coffeeTable.IsFoldable())
	} else {
		fmt.Println("INFORMA hasn't implement coffee table")
	}
}
