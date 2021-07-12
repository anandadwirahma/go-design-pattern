package adapter

import "fmt"

type UsbMicro struct {
	Device string
}

func (m UsbMicro) Charging() {
	fmt.Printf("Device %s is charging using usb micro.\n", m.Device)
}
