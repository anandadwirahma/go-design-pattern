package adapter

import "fmt"

type Client struct {
}

func (c Client) InsertCharger(charger ChargerAdapter) {
	fmt.Println("Client is insert charger.")
	charger.Charging()
}
