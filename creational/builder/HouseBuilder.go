package builder

type houseBuilder struct {
	window string
	wall   string
	door   string
	floor  int
}

func GetHouseBuilder() houseBuilder {
	return houseBuilder{}
}

func (h *houseBuilder) BuildHouse() house {
	return house{
		door:   h.door,
		wall:   h.wall,
		floor:  h.floor,
		window: h.window,
	}
}

func (h *houseBuilder) buildWindow() *houseBuilder {
	h.window = "Window"
	return h
}

func (h *houseBuilder) buildWall() *houseBuilder {
	h.wall = "Wall"
	return h
}

func (h *houseBuilder) buildDoor() *houseBuilder {
	h.door = "Door"
	return h
}

func (h *houseBuilder) buildFloor() *houseBuilder {
	h.floor = 3
	return h
}
