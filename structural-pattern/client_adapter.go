package main

import (
	"fmt"
	"go-design-pattern/structural-pattern/adapter"
)

func CallAdapterMethod() {
	client := adapter.Client{}
	wireless := adapter.Wireles{}

	usbC := adapter.UsbC{Device: "Samsung"}
	client.InsertCharger(usbC)
	usbCToWirelessAdapter := adapter.UsbCToWirelessAdapter{WirelessAdapter: wireless}
	client.InsertCharger(usbCToWirelessAdapter)

	fmt.Println("")

	usbMicro := adapter.UsbMicro{Device: "Xiaomi"}
	client.InsertCharger(usbMicro)
	usbMicroToWirelessAdapter := adapter.UsbMicroToWirelessAdapter{WirelessAdapter: wireless}
	client.InsertCharger(usbMicroToWirelessAdapter)
}
