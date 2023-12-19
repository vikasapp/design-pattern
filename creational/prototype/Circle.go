package prototype

import "fmt"

type Circle struct {
	ID            ShapeType
	Radius        int
	Diameter      int
	Circumference int
}

func NewCircle(radius, diameter, circumference int) Circle {
	return Circle{TypeCircle, radius, diameter, circumference}
}

func (c Circle) GetID() ShapeType {
	return c.ID
}

func (c Circle) PrintPrototype() {
	fmt.Println("Circle Properties are ID:", c.ID, " Radius:", c.Radius, " Diameter:", c.Diameter, " Circumfrence:", c.Circumference)
}

func (c Circle) Clone() IShape {
	return NewCircle(c.Radius, c.Diameter, c.Circumference)
}
