package abstractfactory

type Nike struct {
}

func (a *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			Logo: "Nike",
			Size: 14,
		},
	}
}

func (a *Nike) MakeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			Logo: "Nike",
			Size: 14,
		},
	}
}
