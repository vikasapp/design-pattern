package momento

import "fmt"

type Distance struct {
	value int
}

func NewDistance(distance int) *Distance {
	return &Distance{value: distance}
}

func (d *Distance) Travelled(distance int) {
	d.value -= distance
}

func (d *Distance) Print() {
	fmt.Println("Value of distance is : ", d.value)
}
