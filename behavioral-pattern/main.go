package main

import "fmt"

func main() {
	fmt.Println("========== Chan Of Responsibilty Pattern ==========")
	CallChainOfResponsibility()

	fmt.Println("========== Command Pattern ==========")
	CallCommandMethod()

	fmt.Println("========== Iterator Pattern ==========")
	CallIteratorMethod()
}
