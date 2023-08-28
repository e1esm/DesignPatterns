package main

import "fmt"

type AreaCalculator struct {
	area int
}

func (ac *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
}

func (ac *AreaCalculator) visitForCircle(c *Circle) {
	fmt.Println("Calculating area for circle")
}

func (ac *AreaCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
