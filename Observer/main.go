package main

func main() {
	shirtItem := NewItem("Nike Shirt")

	observerFirst := &Customer{id: "xyz@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)
	shirtItem.updateAvailability()
}
