package composite

import "fmt"

type Product struct {
	Name      string
	Price     int64
	NumOfItem int64
}

func (p Product) GetPrice() (totalPrice int64) {
	fmt.Printf("Item Name: %s\n", p.Name)
	fmt.Printf("Price: %d\n", p.Price)
	fmt.Printf("Qty: %d\n", p.NumOfItem)
	fmt.Println()

	totalPrice = p.Price * p.NumOfItem
	return
}
