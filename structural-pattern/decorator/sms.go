package decorator

import "fmt"

type SMS struct {
	PhoneNumbers []string
}

func (s SMS) SendMessage(message string) {
	for _, phone := range s.PhoneNumbers {
		fmt.Println("send message " + message + " to " + phone + " using SMS.")
	}
}
