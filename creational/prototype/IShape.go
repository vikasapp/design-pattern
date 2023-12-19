package prototype

type ShapeType int

const (
	TypeCircle ShapeType = iota + 1
	TypeSquare
)

type IShape interface {
	GetID() ShapeType
	PrintPrototype()
	Clone() IShape
}
