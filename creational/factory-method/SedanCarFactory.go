package factorymethod

type SedanCarFactory struct {
}

func (s SedanCarFactory) CreateCar(name string) ICar {
	return SedanCar{Name: name}
}
