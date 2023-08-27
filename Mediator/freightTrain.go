package main

import "fmt"

type FreightTrain struct {
	mediator Mediator
}

func (f *FreightTrain) arrive() {
	if !f.mediator.canArrive(f) {
		fmt.Println("FreightTrain: Arrival is blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: arrived")
}

func (f *FreightTrain) depart() {
	fmt.Println("FreightTrain: leaving")
	f.mediator.notifyAboutDeparture()
}

func (f *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted, arriving")
	f.arrive()
}
