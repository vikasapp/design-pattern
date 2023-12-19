package state

type state struct {
	CurrentState ITV
}

func GetStateContext() *state {
	return &state{
		CurrentState: &Off{},
	}
}

func (s *state) SetState(state ITV) {
	s.CurrentState = state
}

func (s *state) GetState() ITV {
	return s.CurrentState
}

func TestStatePattern() {
	state := GetStateContext()
	state.GetState().State()
	state.SetState(&On{})
	state.GetState().State()
}
