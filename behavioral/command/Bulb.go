package command

import "fmt"

type Bulb struct{}

func (b *Bulb) On() {
	fmt.Println("Bulb is on...")
}

func (b *Bulb) Off() {
	fmt.Println("Bulb is off...")
}
