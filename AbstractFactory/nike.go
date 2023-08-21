package main

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe{
			size: 14,
			logo: "nike",
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt{
			logo: "nike",
			size: 14,
		},
	}
}
