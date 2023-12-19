package state

import "fmt"

type Off struct{}

func (o *Off) State() {
	fmt.Println("TV OF...")
}
