package factorymethod

type ICarFactory interface {
	CreateCar(name string) ICar
}

func GetCarFactory(factory ICarFactory, name string) {
	factory.CreateCar(name).GetCar()
}
