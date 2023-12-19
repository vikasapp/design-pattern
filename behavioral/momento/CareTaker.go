package momento

type CareTaker struct {
	momentos []DistanceMomento
}

func NewCareTaker() *CareTaker {
	return &CareTaker{
		momentos: make([]DistanceMomento, 0),
	}
}

func (c *CareTaker) AddMomento(dm DistanceMomento) {
	c.momentos = append(c.momentos, dm)
}

func (c *CareTaker) GetLastIndexMomento() DistanceMomento {
	lastIndex := len(c.momentos) - 1
	lastMomento := c.momentos[lastIndex]
	c.momentos = c.momentos[:lastIndex]
	return lastMomento
}
