package main

type Director struct {
	builder IBuilder
}

func NewDirector(builder IBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) setBuilder(builder IBuilder) {
	d.builder = builder
}

func (d *Director) BuildHouse() *House {
	return d.builder.
		setWindowType().
		setDoorType().
		setNumFloor().
		GetHouse()
}
