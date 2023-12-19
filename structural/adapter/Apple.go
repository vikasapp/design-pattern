package adapter

import "fmt"

type Apple struct{}

func (a *Apple) ChargeAppleMobile() {
	fmt.Println("Apple Mobile Charging...")
}
