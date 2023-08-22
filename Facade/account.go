package main

import "fmt"

type Account struct {
	name string
}

func NewAccount(name string) *Account {
	return &Account{name: name}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name == accountName {
		fmt.Println("Account's verified")
		return nil
	}
	return fmt.Errorf("account name is incorrect")
}
