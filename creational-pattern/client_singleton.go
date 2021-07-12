package main

import (
	"fmt"
	"go-design-pattern/creational-pattern/singleton"
)

func CallSingletonWithLock() {
	for i := 0; i < 10; i++ {
		go singleton.GetDBInstanceWithLock()
	}
	fmt.Scanln()
}

func CallSingletonWithOnce() {
	for i := 0; i < 10; i++ {
		go singleton.GetDBInstanceWithOnce()
	}
	fmt.Scanln()
}
