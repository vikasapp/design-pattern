package observer

type IPublisher interface {
	Regisrty(sub Subscriber)
	NotifyAll(msg string)
}
