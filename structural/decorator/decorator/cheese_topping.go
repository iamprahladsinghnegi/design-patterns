package decorator

type CheeseTopping struct {
	Pizza Pizza
}

func (t *CheeseTopping) GetPrice() int {
	pizzaPrice := t.Pizza.GetPrice()
	return pizzaPrice + 10
}
