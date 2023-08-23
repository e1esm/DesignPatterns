package main

type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func NewGame() *Game {
	return &Game{
		terrorists:        make([]*Player, 0),
		counterTerrorists: make([]*Player, 0),
	}
}

func (g *Game) addTerrorist(dressType string) {
	player := NewPlayer("T", dressType)
	g.terrorists = append(g.terrorists, player)
}

func (g *Game) addCounterTerrorist(dressType string) {
	player := NewPlayer("CT", dressType)
	g.terrorists = append(g.terrorists, player)
}
