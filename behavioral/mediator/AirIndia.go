package mediator

import "fmt"

type AirIndia struct {
	mediator IMediator
}

func (f *AirIndia) requestLanding() {
	if f.mediator.canLand(f) {
		fmt.Println("Air India flight is landing")
	} else {
		fmt.Println("Air India flight is waiting to land")
	}
}

func (f *AirIndia) landed() {
	fmt.Println("Air India flight has landed")
	f.mediator.notify()
}

func (f *AirIndia) permitLanding() {
	fmt.Println("Air India flight has been permitted to landing")
}
