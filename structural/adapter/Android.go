package adapter

import "fmt"

type Android struct{}

func (a *Android) ChargeAndroidMobile() {
	fmt.Println("Android Mobile Charging...")
}
