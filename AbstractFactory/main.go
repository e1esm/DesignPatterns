package main

import "fmt"

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(shoe IShoe) {
	fmt.Printf("Logo: %s\n", shoe.getLogo())
	fmt.Printf("Size: %d\n", shoe.getSize())
}

func printShirtDetails(shirt IShirt) {
	fmt.Printf("Logo: %s\n", shirt.getLogo())
	fmt.Printf("Size: %d\n", shirt.getSize())
}
