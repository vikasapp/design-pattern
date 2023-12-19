package builder

type Director struct {
	builder houseBuilder
}

func NewDirector(hb houseBuilder) *Director {
	return &Director{
		builder: hb,
	}
}

func (d *Director) BuildHouse() house {
	d.builder.buildDoor().buildWall().buildWindow().buildFloor()
	return d.builder.BuildHouse()
}
