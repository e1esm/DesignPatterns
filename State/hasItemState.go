package main

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (h *HasItemState) addItem(count int) error {
	fmt.Printf("%d items were added", count)
	h.vendingMachine.incrementItemCount(count)
	return nil
}

func (h *HasItemState) requestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		return fmt.Errorf("no items presented")
	}
	fmt.Println("Item was requested")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h *HasItemState) insertMoney(i int) error {
	return fmt.Errorf("please select item first")
}

func (h *HasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}
