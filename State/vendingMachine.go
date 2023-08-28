package main

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	itemCount     int
	itemPrice     int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}
	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	v.itemRequested = itemRequestedState
	return v
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}
