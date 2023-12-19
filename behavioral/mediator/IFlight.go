package mediator

type IFlight interface {
	requestLanding()
	landed()
	permitLanding()
}
