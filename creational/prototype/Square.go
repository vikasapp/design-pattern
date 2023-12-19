package prototype

import "fmt"

type Square struct {
	ID      ShapeType
	Length  int
	Breadth int
}

func NewSquare(length, breadth int) Square {
	return Square{TypeSquare, length, breadth}
}

func (s Square) GetID() ShapeType {
	return s.ID
}

func (s Square) PrintPrototype() {
	fmt.Println("Square Properties are ID:", s.ID, " Length:", s.Length, " Breadth:", s.Breadth)
}

func (s Square) Clone() IShape {
	return NewSquare(s.Length, s.Breadth)
}
