package main

import (
	"fmt"

	"go-design-pattern/creational-pattern/factory"
	"go-design-pattern/creational-pattern/factory/logistic/road_logistic"
	"go-design-pattern/creational-pattern/factory/logistic/sea_logistic"
)

func CallFactoryMethod() {
	var logistic factory.Logistic

	fmt.Println("-- Road Logistic --")
	logistic = road_logistic.RoadLogistic{
		Length: 10,
		Width:  5,
		Height: 10,
	}

	transport := logistic.CreateTransport()
	fmt.Printf("Transport: %s\n", transport.GetTransport())

	fmt.Println("-- Sea Logistic --")
	logistic = sea_logistic.SeaLogistic{}

	transport = logistic.CreateTransport()
	fmt.Printf("Transport: %s\n", transport.GetTransport())
}
