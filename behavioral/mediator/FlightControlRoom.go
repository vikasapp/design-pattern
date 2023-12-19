package mediator

type FlightControlRoom struct {
	isRunWayFree bool
	flightQueue  []IFlight
}

func (f *FlightControlRoom) canLand(flight IFlight) bool {
	if f.isRunWayFree {
		f.isRunWayFree = false
		return true
	}
	f.flightQueue = append(f.flightQueue, flight)
	return false
}

func (f *FlightControlRoom) notify() {
	if !f.isRunWayFree {
		f.isRunWayFree = true
	}

	if len(f.flightQueue) > 0 {
		firstFlight := f.flightQueue[0]
		f.flightQueue = f.flightQueue[1:]
		firstFlight.permitLanding()
	}
}
