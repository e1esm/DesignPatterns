package main

type IBuilder interface {
	setWindowType() IBuilder
	setDoorType() IBuilder
	setNumFloor() IBuilder
	GetHouse() *House
}
