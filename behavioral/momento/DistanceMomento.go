package momento

type DistanceMomento struct {
	distance Distance
}

func (dm *DistanceMomento) restore() Distance {
	return dm.distance
}
