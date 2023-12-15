package visitor

type Rectangle struct {
	L int
	B int
}

func (t *Rectangle) Accept(v Visitor) {
	v.visitForRectangle(t)
}

func (t *Rectangle) GetType() string {
	return "rectangle"
}
