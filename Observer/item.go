package main

import "fmt"

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is in stock", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) deregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromSlice(observeList []Observer, toBeRemoved Observer) []Observer {
	for i, v := range observeList {
		if v.getID() == toBeRemoved.getID() {
			observeList[len(observeList)-1], observeList[i] = observeList[i], observeList[len(observeList)-1]
			return observeList[:len(observeList)-1]
		}
	}
	return observeList
}
