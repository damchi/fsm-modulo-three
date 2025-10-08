package fsm

import (
	"fmt"
	"fsm-modulo-three/internal/core/ports"
)

// Ensure service implements the port
var _ ports.ModuloService = (*ModuloService)(nil)

type ModuloService struct{}

// Compute calculates remainder of a binary string modulo `mod` using FSM
func (s *ModuloService) Compute(binary string, mod int) (int, error) {
	if mod <= 0 {
		return 0, fmt.Errorf("mod must be > 0")
	}

	moduloFsm := BuildModuloFSM(mod)
	seq := []rune(binary)

	final, err := moduloFsm.Run(seq)
	if err != nil {
		return 0, err
	}
	return final, nil
}

// BuildModuloFSM generates FSM for any modulo
func BuildModuloFSM(mod int) *FSM {
	// get the state dynamically from the mod
	states := make([]int, mod)
	finals := make([]int, mod)
	for i := 0; i < mod; i++ {
		states[i] = i
		finals[i] = i
	}

	alphabet := []rune{'0', '1'}
	transitions := make(map[int]map[rune]int)
	for _, state := range states {
		transitions[state] = make(map[rune]int)
		for _, b := range alphabet {
			v := int(b - '0')
			transitions[state][b] = (state*2 + v) % mod
		}
	}
	return &FSM{
		States:      states,
		Alphabet:    alphabet,
		StartState:  0,
		FinalStates: finals,
		Transitions: transitions,
	}
}
