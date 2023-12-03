package main

import (
	adapter "github.com/iamprahladsinghnegi/design-patterns/structural/adapter/adapter"
)

func main() {

	client := &adapter.Client{}
	mac := &adapter.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &adapter.Windows{}
	windowsMachineAdapter := &adapter.WindowsAdapter{
		WindowsMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

}
