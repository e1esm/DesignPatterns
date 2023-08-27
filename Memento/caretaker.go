package main

type CareTaker struct {
	mementoArray []*Memento
}

func (ct *CareTaker) addMemento(m *Memento) {
	ct.mementoArray = append(ct.mementoArray, m)
}

func (ct *CareTaker) getMemento(index int) *Memento {
	return ct.mementoArray[index]
}
