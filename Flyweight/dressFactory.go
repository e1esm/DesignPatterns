package main

import "fmt"

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (df *DressFactory) getDressByType(dressType string) (Dress, error) {
	if df.dressMap[dressType] != nil {
		return df.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		//df.dressMap[dressType]
		return df.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {

		return df.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("wrong dress type")
}

func GetDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
