package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewIglooBuilder() IBuilder {
	return &IglooBuilder{}
}

func (ib *IglooBuilder) setWindowType() IBuilder {
	ib.windowType = "Snow Window"
	return ib
}

func (ib *IglooBuilder) setDoorType() IBuilder {
	ib.doorType = "Snow door"
	return ib
}

func (ib *IglooBuilder) setNumFloor() IBuilder {
	ib.floor = 1
	return ib
}

func (ib *IglooBuilder) GetHouse() *House {
	return &House{
		windowType: ib.windowType,
		doorType:   ib.doorType,
		floor:      ib.floor,
	}
}
