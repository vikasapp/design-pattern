package chainofresponsibility

import "fmt"

type OpenFlap struct {
	Next IStep
}

func (o *OpenFlap) Execute() {
	fmt.Println("Opening Laptop")
	o.Next.Execute()
}

func (o *OpenFlap) SetNextStep(step IStep) IStep {
	o.Next = step
	return o.Next
}
