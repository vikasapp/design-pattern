package momento

type DistanceOriginator struct {
	distance Distance
}

func NewDistanceOriginator(distance Distance) *DistanceOriginator {
	return &DistanceOriginator{
		distance: distance,
	}
}

func (do *DistanceOriginator) SaveState() *DistanceMomento {
	return &DistanceMomento{
		distance: do.distance,
	}
}

func (do *DistanceOriginator) GetState() Distance {
	return do.distance
}

func (do *DistanceOriginator) SetState(distance Distance) {
	do.distance = distance
}

func (do *DistanceOriginator) RestoreFromMomento(dm *DistanceMomento) {
	do.distance = dm.restore()
}
