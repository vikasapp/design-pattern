package observer

type ISubscriber interface {
	ReactToPublisher(msg string)
}
