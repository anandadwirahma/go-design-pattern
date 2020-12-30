package main

/*
**notes :
	iGun as a Product Interface
*/

type iGun interface {
	setName(name string)
	setPower(powere int)
	getName() string
	getPower() int
}
