package main

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func NewPlayer(playerType string, dressType string) *Player {
	dress, _ := GetDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		dress:      dress,
		playerType: playerType,
	}
}

func (p *Player) NewLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
