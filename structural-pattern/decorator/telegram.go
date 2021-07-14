package decorator

import "fmt"

type Telegram struct {
	ID string
}

func (t Telegram) SendMessage(message string) {
	fmt.Println("send message " + message + " to " + t.ID + " using Telegram.")
}
