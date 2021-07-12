package main

import (
	"fmt"
)

func main() {
	fmt.Println("========== Factory ==========")
	CallFactoryMethod()

	fmt.Println("\n========== Abstract Factory ==========")
	CallAbstractFactoryMethod()

	fmt.Println("\n========== Builder ==========")
	CallBuilderMethod()

	fmt.Println("\n========== Prototype ==========")
	CallPrototypeMethod()

	fmt.Println("\n========== Singleton With Lock ==========")
	CallSingletonWithLock()

	fmt.Println("\n========== Singleton With Once ==========")
	CallSingletonWithOnce()
}
