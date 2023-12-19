package mediator

type IMediator interface {
	canLand(IFlight) bool
	notify()
}
