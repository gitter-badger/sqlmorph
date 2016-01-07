package parsing

// BaseState provides base functionality to parse states.
type BaseState struct {
	name       string
	nextStates []State
	isEndState bool
}

func (s *BaseState) MakeEndState() {
	s.isEndState = true
}

func (s *BaseState) IsEndState() bool {
	return s.isEndState
}

func (s *BaseState) SetNext(nextStates ...State) {
	s.nextStates = nextStates
}

func (s *BaseState) Next() []State {
	return s.nextStates
}
