package abstractfactory

type Nike struct {
}

func (a *Nike) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			Logo: "Nike",
			Size: 14,
		},
	}
}

func (a *Nike) MakeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			Logo: "Nike",
			Size: 14,
		},
	}
}
