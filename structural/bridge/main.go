package main

import (
	"fmt"

	bridge "github.com/iamprahladsinghnegi/design-patterns/structural/bridge/bridge"
)

func main() {

	hpPrinter := &bridge.Hp{}
	epsonPrinter := &bridge.Epson{}

	macComputer := &bridge.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	windowsComputer := &bridge.Windows{}
	windowsComputer.SetPrinter(hpPrinter)
	windowsComputer.Print()
	fmt.Println()

	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
	fmt.Println()

}
