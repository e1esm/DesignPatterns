package main

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine was already given to this patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medicine's being given to the patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}
