package main

import (
	"fmt"
	"go-design-pattern/structural-pattern/decorator/example2/door"
)

func main() {
	d := door.Door{}
	d.Open()

	fmt.Println("=========")

	ed := door.NewElectronicKeyDoor(&d)
	ed.Open()

	fmt.Println("=========")

	md := door.NewMagicalDoor(ed)
	md.Open()
}
