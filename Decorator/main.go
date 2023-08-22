package main

import "fmt"

func main() {
	pizza := &VeggePizza{}

	pizzaWithTomatoes := &TomatoTopping{pizza: pizza}

	pizzaWithCheese := &CheeseTopping{pizza: pizzaWithTomatoes}

	fmt.Printf("Overall price: %d\n", pizzaWithCheese.getPrice())
}
