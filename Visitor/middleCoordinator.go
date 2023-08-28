package main

import "fmt"

type MiddleCoordinates struct {
	x int
	y int
}

func (mc *MiddleCoordinates) visitForSquare(s *Square) {
	fmt.Println("Calculating middle point for square")
}

func (mc *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point for circle")
}

func (mc *MiddleCoordinates) visitForRectangle(r *Rectangle) {
	fmt.Println("Calculating middle point for rectangle")
}
