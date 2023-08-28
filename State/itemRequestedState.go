package main

import "fmt"

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) addItem(i2 int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("item's already requested")
}

func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money isn't enough")
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}
