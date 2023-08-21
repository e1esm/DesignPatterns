package main

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewNormalBuilder() IBuilder {
	return &NormalBuilder{}
}

func (nb *NormalBuilder) setWindowType() IBuilder {
	nb.windowType = "Wooden Window"
	return nb
}

func (nb *NormalBuilder) setDoorType() IBuilder {
	nb.doorType = "Wooden door"
	return nb
}

func (nb *NormalBuilder) setNumFloor() IBuilder {
	nb.floor = 2
	return nb
}

func (nb *NormalBuilder) GetHouse() *House {
	return &House{
		windowType: nb.windowType,
		doorType:   nb.doorType,
		floor:      nb.floor,
	}
}
