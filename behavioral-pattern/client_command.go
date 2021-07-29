package main

import "go-design-pattern/behavioral-pattern/command"

func CallCommandMethod() {
	tv := &command.TV{}

	onCommand := &command.OnCommand{
		Device: tv,
	}

	offCommand := &command.OffCommand{
		Device: tv,
	}

	onButton := &command.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &command.Button{
		Command: offCommand,
	}
	offButton.Press()
}
