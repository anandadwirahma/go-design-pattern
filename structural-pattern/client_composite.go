package main

import (
	"fmt"
	"go-design-pattern/structural-pattern/composite"
)

func CallCompositeMethod() {
	product1 := &composite.Product{Name: "Susu", Price: 5000, NumOfItem: 2}
	product2 := composite.Product{Name: "Indomie", Price: 3000, NumOfItem: 2}
	product3 := composite.Product{Name: "Beras", Price: 10000, NumOfItem: 1}

	box1 := &composite.Boxes{Name: "Box1"}
	box1.Add(product1)
	box1.Add(product3)
	box3 := &composite.Boxes{Name: "Box3"}
	box3.Add(box1)

	box2 := composite.Boxes{Name: "Box2"}
	box2.Add(product2)

	box2.Add(box3)

	totalPrice := box2.GetPrice()

	fmt.Printf("Total Price = %d", totalPrice)
}
