package builder

type NormalBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) SetWindowType() {
	b.WindowType = "Wooden window"
}

func (b *NormalBuilder) SetDoorType() {
	b.DoorType = "Wooden door"
}

func (b *NormalBuilder) SetNumFloor() {
	b.Floor = 2
}

func (b *NormalBuilder) GetHouse() House {
	return House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
