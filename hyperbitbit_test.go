package hyperbitbit

import (
	"testing"
)

func TestP(t *testing.T) {
	tests := map[uint64]uint64{
		8: 1,
		7: 3,
		6: 2,
		5: 2,
		4: 1,
		3: 2,
		2: 1,
		1: 1,
		0: 0,
	}
	for val, expected := range tests {
		res := p(val)
		if res != expected {
			t.Errorf("Expected %d bits set to 1 in %d, got %d", expected, val, res)
		}
	}
}

func TestR(t *testing.T) {
	tests := map[uint64]uint64{
		8: 0,
		7: 3,
		6: 0,
		5: 1,
		4: 0,
		3: 2,
		2: 0,
		1: 1,
		0: 0,
	}
	for val, expected := range tests {
		res := r(val)
		if res != expected {
			t.Errorf("Expected %d trailing 1 in %d, got %d", expected, val, res)
		}
	}
}
