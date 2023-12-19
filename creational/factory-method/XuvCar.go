package factorymethod

import "fmt"

type XuvCar struct {
	Name string
}

func (x XuvCar) GetCar() {
	fmt.Println("XUV Car :" + x.Name)
}
