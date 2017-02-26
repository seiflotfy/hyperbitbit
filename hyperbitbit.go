package hyperbitbit

import (
	"math"

	metro "github.com/dgryski/go-metro"
)

func p(sketch uint64) uint64 {
	var count uint64
	for i := uint8(0); i < 64 && sketch > 0; i++ {
		if (sketch>>i)%2 == 1 {
			count++
		}
	}
	return count
}

// Calculate the position of the leftmost 1-bit.
func rho(val uint64, max uint64) (r uint64) {
	for r = 0; val&0x1 == 1 && r <= max; r++ {
		val >>= 1
	}
	return r
}

// HyperBitBit sketch for cardinality estimation
type HyperBitBit struct {
	lgN     uint8
	sketch  uint64
	sketch2 uint64
}

// New creates a new HyperBitBit sketch
func New() *HyperBitBit {
	return &HyperBitBit{
		lgN:     5,
		sketch:  0,
		sketch2: 0,
	}
}

// Add a value ([]byte) to the sketch
func (hbb *HyperBitBit) Add(val []byte) {
	x := metro.Hash64([]byte(val), 42)
	k := x << 58 >> 58
	r := rho(x>>6, 58)

	if r > uint64(hbb.lgN) {
		hbb.sketch = hbb.sketch | (uint64(1) << k)
	}
	if r > uint64(hbb.lgN+1) {
		hbb.sketch2 = hbb.sketch2 | (uint64(1) << k)
	}
	if p(hbb.sketch) > 31 {
		hbb.sketch = hbb.sketch2
		hbb.sketch2 = 0
		hbb.lgN++
	}
}

// Cardinality returns the estimated number of unique elements added to the sketch
func (hbb *HyperBitBit) Cardinality() uint64 {
	lgN := float64(hbb.lgN)
	p := float64(p(hbb.sketch))
	return uint64(math.Pow(2, lgN+5.4+p/32.0))
}
