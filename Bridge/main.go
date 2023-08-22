package main

func main() {
	hpPrinter := &HP{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()

	windowsComputer := &Windows{}
	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
}
