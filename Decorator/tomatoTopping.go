package main

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice() + 7
	return pizzaPrice
}
