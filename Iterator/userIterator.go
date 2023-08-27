package main

type UserIterator struct {
	index int
	Users []*User
}

func (ui *UserIterator) Next() *User {
	if ui.HasNext() {
		user := ui.Users[ui.index]
		ui.index++
		return user
	}
	return nil
}

func (ui *UserIterator) HasNext() bool {
	if ui.index < len(ui.Users) {
		return true
	}
	return false
}
