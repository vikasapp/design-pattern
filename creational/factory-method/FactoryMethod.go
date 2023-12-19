package factorymethod

func TestFactoryMethod() {
	GetCarFactory(SedanCarFactory{}, "Honda")
	GetCarFactory(XuvCarFactory{}, "Mahindra")
	GetCarFactory(HachbackCarFactory{}, "Maruti")
}
