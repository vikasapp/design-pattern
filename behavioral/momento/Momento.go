package momento

func TestMomentoPattern() {
	dis := NewDistance(100)
	careTaker := NewCareTaker()
	originator := NewDistanceOriginator(*dis)
	dis.Print()

	// First Momento
	m1 := originator.SaveState()
	careTaker.AddMomento(*m1)

	dis2 := originator.GetState()
	dis2.Print()
	dis2.Travelled(50)
	dis2.Print()

	// Cecon Momento
	originator.SetState(dis2)
	m2 := originator.SaveState()
	careTaker.AddMomento(*m2)

	dis2.Travelled(25)
	dis2.Print()

	// Restore
	lastMomento := careTaker.GetLastIndexMomento()
	originator.RestoreFromMomento(&lastMomento)
	d3 := originator.GetState()
	d3.Print() // Restore
	lastMomento2 := careTaker.GetLastIndexMomento()
	originator.RestoreFromMomento(&lastMomento2)
	d4 := originator.GetState()
	d4.Print()
}
