package cor

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment done")
	}

	fmt.Println("Cashier getting money from patient " + p.Name)
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
