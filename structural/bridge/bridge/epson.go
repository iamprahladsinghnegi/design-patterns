package bridge

import "fmt"

type Epson struct{}

func (h *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
