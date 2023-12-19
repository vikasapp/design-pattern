package observer

import "fmt"

type Subscriber struct {
	ID int
}

func GetNewSubscriber() Subscriber {
	return Subscriber{}
}

func (s Subscriber) ReactToPublisher(msg string) {
	fmt.Println("Subscriber:", s.ID, " receive publisher msg:"+msg)
}
