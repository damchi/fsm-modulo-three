package fsm

import (
	"testing"
)

func TestFSM_Run(t *testing.T) {
	tests := []struct {
		mod   int
		input string
	}{
		{3, "1010"},    // 10 % 3 = 1
		{3, "111"},     // 7 % 3 = 1
		{2, "1101"},    // 13 % 2 = 1
		{5, "11111"},   // 31 % 5 = 1
		{7, "1010101"}, // 85 % 7 = 1
		{4, ""},        // empty input = 0
	}

	for _, tt := range tests {
		fsm := BuildModuloFSM(tt.mod)
		seq := []rune(tt.input)
		final, err := fsm.Run(seq)
		if err != nil {
			t.Errorf("FSM.Run(%s) returned error: %v", tt.input, err)
		}

		// Compute expected remainder
		expected := 0
		if tt.input != "" {
			n := 0
			for _, c := range tt.input {
				if c != '0' && c != '1' {
					t.Fatalf("invalid binary character: %c", c)
				}
				v := int(c - '0')
				n = n*2 + v
			}
			expected = n % tt.mod
		}

		if final != expected {
			t.Errorf("FSM.Run(%s) = %d; want %d", tt.input, final, expected)
		}
	}
}

func TestFSM_InvalidSymbol(t *testing.T) {
	fsm := BuildModuloFSM(3)
	err := fsm.Input('2')
	if err == nil {
		t.Errorf("expected error for invalid symbol, got nil")
	}
}
