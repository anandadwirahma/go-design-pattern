package main

import "go-design-pattern/structural-pattern/decorator"

func CallDecoratorMethod() {
	var notifier decorator.Notifier

	sms := decorator.SMS{
		PhoneNumbers: []string{"082111222", "082333444", "082555666"},
	}
	telegram := decorator.Telegram{
		ID: "@wongndeso",
	}

	messenger := decorator.Messenger{
		Notifier: []decorator.Notifier{sms, telegram},
	}
	notifier = messenger
	notifier.SendMessage("hallo indonesia")
}
