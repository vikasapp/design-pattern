package factorymethod

import "fmt"

type SedanCar struct {
	Name string
}

func (s SedanCar) GetCar() {
	fmt.Println("Sedan Car :" + s.Name)
}
