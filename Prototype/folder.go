package main

import "fmt"

type Folder struct {
	Children []INode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, v := range f.Children {
		v.print(indentation + indentation)
	}
}

func (f *Folder) clone() INode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []INode
	for _, v := range f.Children {
		clonedCopy := v.clone()
		tempChildren = append(tempChildren, clonedCopy)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}
