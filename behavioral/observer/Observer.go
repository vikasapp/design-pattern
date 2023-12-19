package observer

func TestObserverPattern() {
	pub := GetNewPublisher()
	for i := 1; i <= 20; i++ {
		sub := GetNewSubscriber()
		sub.ID = i
		pub.Regisrty(sub)
	}
	pub.NotifyAll("OK")
}
