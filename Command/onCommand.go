package main

type onCommand struct {
	device Device
}

func (on *onCommand) execute() {
	on.device.on()
}
