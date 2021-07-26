package door

import (
	"fmt"
	"math/rand"
)

//type BaseDecorator struct {
//	Door Opener
//}
//
//func (b BaseDecorator) Open() {
//	fmt.Println("Base decorator open")
//	b.Door.Open()
//}

func NewElectronicKeyDoor(opener Opener) Opener {
	return &ElectronicKey{
		opener: opener,
	}
}

type ElectronicKey struct {
	opener Opener
}

func (e ElectronicKey) Open() {
	fmt.Println("this is electornic key")
	e.ConnectToWifi()
	e.opener.Open()
}

func (e ElectronicKey) ConnectToWifi() {
	fmt.Println("connecting to wifi")
}

func NewMagicalDoor(opener Opener) Opener {
	return &MagicalKey{
		opener: opener,
	}
}

type MagicalKey struct {
	opener Opener
}

func (e MagicalKey) Open() {
	fmt.Println("magical key is dangerous")
	if e.CanWeUseMagic() {
		fmt.Println("magic is working")
	} else {
		fmt.Println("can't we use magic, use your hand")
	}
	e.opener.Open()
}

func (e MagicalKey) CanWeUseMagic() bool {
	return rand.Int() > 0
}
