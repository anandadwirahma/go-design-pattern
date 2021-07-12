package bridge

import "fmt"

type Shape interface {
	Draw()
	Area() float64
}

type Circle struct {
	Color  Color
	Radius int
}

func (c Circle) Draw() {
	if c.Color != nil {
		fmt.Println("Drawing circle with color is " + c.Color.Hex())
	} else {
		fmt.Println("Drawing circle")
	}
}

func (c Circle) Area() float64 {
	return 3.14 * float64(c.Radius) * float64(c.Radius)
}

type Square struct {
	Color  Color
	Length int
}

func (s Square) Clone() Square {
	return Square{
		Color:  s.Color,
		Length: s.Length,
	}
}

func (s Square) Draw() {
	if s.Color != nil {
		fmt.Println("Drawing square with color is " + s.Color.Hex())
	} else {
		fmt.Println("Drawing square")
	}
}

func (s Square) Area() float64 {
	area := s.Length * s.Length
	return float64(area)
}
