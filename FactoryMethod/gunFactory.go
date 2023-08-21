package main

import "errors"

func getGun(gunName string) (IGun, error) {
	switch {
	case gunName == "AK-47":
		return NewAK47(), nil
	case gunName == "Musket":
		return NewMusket(), nil
	default:
		return nil, errors.New("invalid gun name")
	}
}
