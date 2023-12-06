package main

import (
	"fmt"

	decorator "github.com/iamprahladsinghnegi/design-patterns/structural/decorator/decorator"
)

func main() {

	pizza := &decorator.VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &decorator.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
