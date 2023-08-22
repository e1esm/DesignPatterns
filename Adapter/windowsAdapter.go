package main

import "fmt"

type WindowsAdapter struct {
	WindowsMachine *Windows
}

func (wa *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts lightning to the USB")
	wa.WindowsMachine.InsertIntoUSBPort()
}
