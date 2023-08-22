package main

import "fmt"

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Println("Searching for keyword", keyword)
}

func (f *File) getName() string {
	return f.name
}
