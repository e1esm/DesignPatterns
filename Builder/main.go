package main

import "fmt"

func main() {
	normalBuilder := NewNormalBuilder()
	iglooBuilder := NewIglooBuilder()

	director := NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("%v\n", *normalHouse)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()
	fmt.Printf("%v\n", *iglooHouse)
}
