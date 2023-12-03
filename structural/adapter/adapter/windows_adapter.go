package adapter

type WindowsAdapter struct {
	WindowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	println("Adapter converts Lightning signal to USB.")
	w.WindowsMachine.InsertIntoUSBPort()
}
