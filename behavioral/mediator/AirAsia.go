package mediator

import "fmt"

type AirAsia struct {
	mediator IMediator
}

func (f *AirAsia) requestLanding() {
	if f.mediator.canLand(f) {
		fmt.Println("Air Asia flight is landing")
	} else {
		fmt.Println("Air Asia flight is waiting to land")
	}
}

func (f *AirAsia) landed() {
	fmt.Println("Air Asia flight has landed")
	f.mediator.notify()
}

func (f *AirAsia) permitLanding() {
	fmt.Println("Air Asia flight has been permitted to landing")
}
