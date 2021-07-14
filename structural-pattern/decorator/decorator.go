package decorator

type Messenger struct {
	Notifier []Notifier
}

func (m Messenger) SendMessage(message string) {
	for _, notifier := range m.Notifier {
		notifier.SendMessage(message)
	}
}
