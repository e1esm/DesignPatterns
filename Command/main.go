package main

func main() {
	tv := &TV{}

	onCommand := &onCommand{
		tv,
	}
	offCommand := &offCommand{
		tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	offButton := &Button{
		command: offCommand,
	}
	onButton.press()
	offButton.press()
}
