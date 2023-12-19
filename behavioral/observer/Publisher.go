package observer

import "fmt"

type Publisher struct {
	SubscriberList []ISubscriber
}

func GetNewPublisher() Publisher {
	return Publisher{SubscriberList: make([]ISubscriber, 0)}
}

func (p *Publisher) Regisrty(sub ISubscriber) {
	p.SubscriberList = append(p.SubscriberList, sub)
}

func (p *Publisher) NotifyAll(msg string) {
	for _, sub := range p.SubscriberList {
		fmt.Println("Publisher notifying subscriber with ID:", sub.(Subscriber).ID)
		sub.ReactToPublisher(msg)
	}
}
