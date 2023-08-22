package main

import "fmt"

type HP struct {
}

func (hp *HP) PrintFile() {
	fmt.Println("Printing by a HP printer")
}
