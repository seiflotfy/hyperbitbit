package hyperbitbit

import (
	"math"
	"math/rand"
	"testing"
	"time"
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
		res := rho(val, 58)
		if res != expected {
			t.Errorf("Expected %d trailing 1 in %d, got %d", expected, val, res)
		}
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n uint32) string {
	b := make([]byte, n)
	for i := uint32(0); i < n; i++ {
		b[i] = letterBytes[rand.Int()%len(letterBytes)]
	}
	return string(b)
}

func TestCardinality(t *testing.T) {
	hbb := New()

	step := 10000
	unique := map[string]bool{}

	for i := 1; len(unique) <= 10000000; i++ {
		str := RandStringBytesMaskImprSrc(rand.Uint32() % 32)
		hbb.Add([]byte(str))
		unique[str] = true

		if len(unique)%step == 0 {
			exact := len(unique)
			step *= 10
			res := int(hbb.Cardinality())
			ratio := 100 * math.Abs(float64(res-exact)) / float64(exact)

			expectedError := 0.1

			if float64(res) < float64(exact)-(float64(exact)*expectedError) || float64(res) > float64(exact)+(float64(exact)*expectedError) {
				t.Errorf("Exact %d, got %d which is %.2f%% error", exact, res, ratio)
			}
		}
	}
}
