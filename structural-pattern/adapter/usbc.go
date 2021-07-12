package adapter

import "fmt"

type UsbC struct {
	Device string
}

func (c UsbC) Charging() {
	fmt.Printf("Device %s is charging using usb C.\n", c.Device)
}
