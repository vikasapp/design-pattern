package factorymethod

import "fmt"

type HachbackCar struct {
	Name string
}

func (h HachbackCar) GetCar() {
	fmt.Println("Hackback Car :" + h.Name)
}
