package main

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment's done")
		return
	}
	fmt.Println("Cashier's getting money from the patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
