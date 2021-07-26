package main

import (
	"fmt"
	"go-design-pattern/structural-pattern/facade/example2/customerservice"
)

func main() {
	response := customerservice.HandleTicket(customerservice.Ticket{
		UserEmail: "ananda",
		Type:      "gofood",
		Message:   "food never deliver",
	})

	fmt.Println("response: ", response)
}
