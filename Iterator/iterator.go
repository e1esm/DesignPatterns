package main

type Iterator interface {
	HasNext() bool
	Next() *User
}
