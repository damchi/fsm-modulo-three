package fsm_test

import (
	"fsm-modulo-three/internal/fsm"
	"testing"
)

func TestModuloService_Compute(t *testing.T) {
	service := &fsm.ModuloService{}

	tests := []struct {
		name     string
		binary   string
		mod      int
		expected int
		wantErr  bool
	}{
		// Normal cases
		{"binary 0 mod 3", "0", 3, 0, false},
		{"binary 1 mod 3", "1", 3, 1, false},
		{"binary 10 mod 3", "10", 3, 2, false},
		{"binary 11 mod 3", "11", 3, 0, false},
		{"binary 1010 mod 3", "1010", 3, 1, false},
		{"binary 1101 mod 7", "1101", 7, 6, false},

		// Edge cases
		{"empty binary string", "", 3, 0, false},  // Empty string should result in start state = 0
		{"binary all zeros", "0000", 5, 0, false}, // All zeros always lead to 0
		{"binary all ones", "1111", 2, 1, false},  // mod 2, all ones alternate states

		// Invalid cases
		{"mod zero", "1010", 0, 0, true},
		{"negative mod", "1010", -3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Compute(tt.binary, tt.mod)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Compute() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.expected {
				t.Errorf("Compute() = %v, want %v", got, tt.expected)
			}
		})
	}
}
