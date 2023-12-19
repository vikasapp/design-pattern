package factorymethod

type XuvCarFactory struct {
}

func (s XuvCarFactory) CreateCar(name string) ICar {
	return XuvCar{Name: name}
}
