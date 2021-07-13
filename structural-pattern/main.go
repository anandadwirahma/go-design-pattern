package main

import "fmt"

func main() {
	fmt.Println("========== Adapter Pattern ==========")
	CallAdapterMethod()

	fmt.Println("\n========== Bridge Pattern ==========")
	CallBridgeMethod()

	fmt.Println("\n========== Composite Pattern ==========")
	CallCompositeMethod()
}
