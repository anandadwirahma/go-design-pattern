package main

import (
	"fmt"

	"go-design-pattern/creational-pattern/prototype/factory"
)

func CallPrototypeMethod() {
	airbus := factory.Plane{
		Name:       "airbus",
		Engine:     "C01AA",
		IsUseRadar: true,
	}
	airbusClone := airbus.Clone()

	boeing := factory.Plane{
		Name:       "boeing",
		Engine:     "B02TT",
		IsUseRadar: false,
	}
	boeingClone := boeing.Clone()

	fmt.Println("\n--Plane Airbus--")
	airbus.Print()
	airbusClone.Print()

	fmt.Println("\n--Plane Boeing--")
	boeing.Print()
	boeingClone.Print()
}
