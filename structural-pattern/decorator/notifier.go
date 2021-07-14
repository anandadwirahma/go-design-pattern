package decorator

type Notifier interface {
	SendMessage(message string)
}
