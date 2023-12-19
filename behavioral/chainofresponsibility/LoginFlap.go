package chainofresponsibility

import "fmt"

type LoginFlap struct {
	Next IStep
}

func (o *LoginFlap) Execute() {
	fmt.Println("Logining Laptop")
	fmt.Println("Flap Opened")
}

func (o *LoginFlap) SetNextStep(step IStep) IStep {
	o.Next = step
	return o.Next
}
