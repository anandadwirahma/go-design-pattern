package adapter

import "fmt"

type ChargerAdapter interface {
	Charging()
}

type UsbCToWirelessAdapter struct {
	WirelessAdapter Wireles
}

func (ca UsbCToWirelessAdapter) Charging() {
	fmt.Println("convert usb c to wireless charging.")
	ca.WirelessAdapter.WirelessCharging()
}

type UsbMicroToWirelessAdapter struct {
	WirelessAdapter Wireles
}

func (ma UsbMicroToWirelessAdapter) Charging() {
	fmt.Println("convert usb micro to wireless charging.")
	ma.WirelessAdapter.WirelessCharging()
}
