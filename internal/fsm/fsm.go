package fsm

import "fmt"

// FSM represents a finite state machine with int states and rune inputs.
type FSM struct {
	States      []int
	Alphabet    []rune
	StartState  int
	FinalStates []int
	Transitions map[int]map[rune]int
	Current     int
}

func (f *FSM) Reset() {
	f.Current = f.StartState
}

func (f *FSM) Input(symbol rune) error {
	if _, ok := f.Transitions[f.Current][symbol]; !ok {
		return fmt.Errorf("no transition from %v on %v", f.Current, symbol)
	}
	f.Current = f.Transitions[f.Current][symbol]
	return nil
}

func (f *FSM) Run(seq []rune) (int, error) {
	f.Reset()
	for _, s := range seq {
		if err := f.Input(s); err != nil {
			return f.Current, err
		}
	}
	return f.Current, nil
}
