package main

import "fmt"

type PassengerTrain struct {
	mediator Mediator
}

func (g *PassengerTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("PassengerTrain: Arrival is blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: arrived")
}

func (g *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.arrive()
}
