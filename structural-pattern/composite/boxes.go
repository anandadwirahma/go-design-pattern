package composite

import "fmt"

type Boxes struct {
	Name  string
	Carts []Cart
}

func (b *Boxes) GetPrice() (price int64) {
	fmt.Println("-- Open " + b.Name)
	for _, c := range b.Carts {
		price = price + c.GetPrice()
	}

	return
}

func (b *Boxes) Add(c Cart) {
	b.Carts = append(b.Carts, c)
}
