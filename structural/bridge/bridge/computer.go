package bridge

type Computer interface {
	Print()
	setPrinter(printer Printer)
}
