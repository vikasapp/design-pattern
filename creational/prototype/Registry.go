package prototype

type Registry struct {
	RegistryList map[ShapeType]IShape
}

func NewRegistry() Registry {
	return Registry{RegistryList: make(map[ShapeType]IShape, 0)}
}

func (r Registry) AddRegistry(shape IShape) {
	r.RegistryList[shape.GetID()] = shape
}

func (r Registry) RemoveRegistry(shapeType ShapeType) {
	delete(r.RegistryList, shapeType)
}
