package main

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (h *HasMoneyState) requestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (h *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (h *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (h *HasMoneyState) dispenseItem() error {
	fmt.Printf("Dispensing item")
	h.vendingMachine.itemCount -= 1
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
	} else {
		h.vendingMachine.setState(h.vendingMachine.hasItem)
	}
	return nil
}
