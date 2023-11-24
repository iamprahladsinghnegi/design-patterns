package abstractfactory

type Adidas struct {
}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			Logo: "Adidas",
			Size: 14,
		},
	}
}
