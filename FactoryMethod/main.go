package main

import "fmt"

func main() {
	ak47, _ := getGun("AK-47")
	musket, _ := getGun("Musket")
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Println(g.getName(), g.getPower())
}
