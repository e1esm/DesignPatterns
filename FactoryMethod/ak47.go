package main

type AK47 struct {
	Gun
}

func NewAK47() IGun {
	return &AK47{
		Gun{
			name:  "AK-47",
			power: 90,
		},
	}
}
