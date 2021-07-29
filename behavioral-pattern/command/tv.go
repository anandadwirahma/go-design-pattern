package command

import "fmt"

type TV struct {
	isRunning bool
}

func (t *TV) On() {
	t.isRunning = true
	fmt.Println("Turning TV on")
}

func (t *TV) Off() {
	t.isRunning = false
	fmt.Println("Turning TV off")
}
