package main

import (
	"fmt"
	"go-design-pattern/structural-pattern/facade"
	"log"
)

func CallFacadeMethod() {
	fmt.Println()
	walletFacade := facade.NewWalletFacade("ananda", 12345)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("ananda", 12345, 1000)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	fmt.Println()
	err = walletFacade.DeductMoneyForWallet("ananda", 12345, 1000)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	fmt.Println()
	err = walletFacade.DeductMoneyForWallet("ananda", 12345, 100)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}
