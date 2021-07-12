package adapter

import "fmt"

type Wireles struct {
}

func (w Wireles) WirelessCharging() {
	fmt.Println("The device is charging using wireless charging.")
}
