package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(30)
	for i := 0; i < 30; i++ {
		go func() {
			defer wg.Done()
			getInstance()
		}()
	}
	wg.Wait()
}
