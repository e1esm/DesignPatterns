package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := NewWalletFacade("abc", 1234)
	fmt.Println()
	err := walletFacade.addMoneyToWallet("abc", 1234, 19)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}
	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}
}
