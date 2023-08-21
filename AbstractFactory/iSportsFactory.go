package main

import (
	"errors"
)

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch {
	case brand == "adidas":
		return &Adidas{}, nil
	case brand == "nike":
		return &Nike{}, nil
	default:
		return nil, errors.New("wrong type of brand")
	}
}
