package factorymethod

type HachbackCarFactory struct {
}

func (s HachbackCarFactory) CreateCar(name string) ICar {
	return HachbackCar{Name: name}
}
