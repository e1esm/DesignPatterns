package main

type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun{
			name:  "Musket",
			power: 50,
		},
	}
}
