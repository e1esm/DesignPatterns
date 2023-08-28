package main

import "log"

func main() {
	vm := NewVendingMachine(10, 1)
	err := vm.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vm.insertMoney(1)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vm.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

}
