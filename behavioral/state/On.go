package state

import "fmt"

type On struct{}

func (o *On) State() {
	fmt.Println("TV ON...")
}
