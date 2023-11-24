package factorymethod

type IGun interface {
	GetName() string
	GetPower() int
	SetName(name string)
	SetPower(power int)
}
