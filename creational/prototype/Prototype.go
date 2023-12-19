package prototype

func LoadRegistry(registry Registry) {
	circle := NewCircle(5, 10, 15)
	square := NewSquare(20, 20)
	registry.AddRegistry(circle)
	registry.AddRegistry(square)
}

func TestPrototyprPattern() {
	registry := NewRegistry()
	LoadRegistry(registry)
	registry.RegistryList[TypeCircle].Clone().PrintPrototype()
	registry.RegistryList[TypeSquare].Clone().PrintPrototype()
}
