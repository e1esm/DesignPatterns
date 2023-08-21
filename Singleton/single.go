package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

type Single struct {
}

var singleInstance *Single

func getInstance() *Single {
	if singleInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if singleInstance == nil {
			fmt.Println("creating singleton object")
			singleInstance = &Single{}
		} else {
			fmt.Println("singleton object is already created")
		}
	} else {
		fmt.Println("Already created")
	}
	return singleInstance
}
