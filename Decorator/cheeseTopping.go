package main

type CheeseTopping struct {
	pizza IPizza
}

func (ct *CheeseTopping) getPrice() int {
	pizzaPrice := ct.getPrice() + 10
	return pizzaPrice
}
