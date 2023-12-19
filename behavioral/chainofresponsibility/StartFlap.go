package chainofresponsibility

import "fmt"

type StartFlap struct {
	Next IStep
}

func (o *StartFlap) Execute() {
	fmt.Println("Starting Laptop")
	o.Next.Execute()
}

func (o *StartFlap) SetNextStep(step IStep) IStep {
	o.Next = step
	return o.Next
}
