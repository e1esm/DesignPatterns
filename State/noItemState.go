package main

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (n *NoItemState) addItem(i int) error {
	n.vendingMachine.incrementitemCount(i)
	n.vendingMachine.setState(n.vendingMachine.hasItem)
	return nil
}

func (n *NoItemState) requestItem() error {
	return fmt.Errorf("item's out of stock")
}

func (n *NoItemState) insertMoney(i int) error {
	return fmt.Errorf("item's out of stock")
}

func (n *NoItemState) dispenseItem() error {
	return fmt.Errorf("item's out of stock")
}
