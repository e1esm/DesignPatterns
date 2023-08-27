package main

type offCommand struct {
	device Device
}

func (oc *offCommand) execute() {
	oc.device.off()
}
