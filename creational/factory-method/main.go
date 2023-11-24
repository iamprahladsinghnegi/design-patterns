package main

import (
	"fmt"

	factorymethod "github.com/iamprahladsinghnegi/design-patterns/creational/factory-method/factory-method"
)

func main() {
	ak47, _ := factorymethod.GetGun("ak47")
	musket, _ := factorymethod.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factorymethod.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
