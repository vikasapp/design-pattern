package mediator

func TestMediatorPattern() {
	fControlRoom := &FlightControlRoom{isRunWayFree: true}
	airIndia := &AirIndia{mediator: fControlRoom}
	airAsia := &AirAsia{mediator: fControlRoom}

	airIndia.requestLanding()
	airAsia.requestLanding()
	airIndia.landed()
}
