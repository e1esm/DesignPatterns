package main

import "fmt"

func main() {
	user1 := &User{
		name: "a",
		age:  19,
	}
	user2 := &User{
		name: "b",
		age:  19,
	}

	uc := UserCollection{
		[]*User{user1, user2},
	}
	iterator := uc.NewIterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}

}
