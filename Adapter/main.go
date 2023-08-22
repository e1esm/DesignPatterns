package main

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsWrapper := &WindowsAdapter{WindowsMachine: windowsMachine}
	windowsWrapper.InsertIntoLightningPort()

	client.InsertLightningConnectorIntoComputer(windowsWrapper)
}
