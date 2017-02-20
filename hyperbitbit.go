package hyperbitbit

import (
	"math"

	metro "github.com/dgryski/go-metro"
)

func p(sketch uint64) uint64 {
	var count uint64
	for i := uint8(0); i < 64; i++ {
		if (sketch>>i)%2 == 1 {
			count++
		}
	}
	return count
}

func r(x uint64) uint64 {
	// Get position of first 0
	r := uint64(0)
	for i := uint8(0); i < 64; i++ {
		if (x>>i)%2 == 0 {
			break
		}
		r++
	}
	return uint64(r)
}

func hash2(x, M uint64) uint64 {
	m := uint64(math.Log(float64(M)) / math.Log(2))
	return x >> (M - m)
}

// HyperBitBit ...
type HyperBitBit struct {
	lgN     uint8
	sketch  uint64
	sketch2 uint64
}

// NewHyperBitBit ...
func NewHyperBitBit(lgN uint8) *HyperBitBit {
	return &HyperBitBit{
		lgN:     5,
		sketch:  0,
		sketch2: 0,
	}
}

// Add ...
func (hbb *HyperBitBit) Add(val []byte) {
	x := metro.Hash64([]byte(val), 42)
	k := hash2(x, 64)
	r := r(x)

	if r > uint64(hbb.lgN) {
		hbb.sketch = hbb.sketch | (1 << k)
	}
	if r > uint64(hbb.lgN+1) {
		hbb.sketch2 = hbb.sketch2 | (1 << k)
	}
	if p(hbb.sketch) > 31 {
		hbb.sketch = hbb.sketch2
		hbb.lgN++
		hbb.sketch2 = 0
	}
}

// Get ...
func (hbb *HyperBitBit) Get() uint64 {
	return uint64(math.Pow(2, float64(hbb.lgN)+5.4+(float64(p(hbb.sketch))/32.0)))
}
