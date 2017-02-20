# HyperBitBit
Initial implementation of `HyperBitBit` as seen on the [presentation](https://www.cs.princeton.edu/~rs/talks/AC11-Cardinality.pdf) by Robert Sedgewick, that aims to beat `HyperLogLog` in practice:

Necessary characteristics of a better algorithm
* Makes one pass through the stream.
* Uses a few dozen machine instructions per value
* Uses a few hundred bits
* Achieves 10% relative accuracy or better

Conjecture. On practical data, HyperBitBit, for N < 2^64,
* Uses 128 + 6 bits. (in this implementation case 128 + 8)
* Estimates cardinality within 10% of the actual.

## Note
This is an implementation that still needs iterations and experimentation. Any help is more than welcome.
* Small cardinalities are way off
* Reinsertings same value increases counter (unlike HyperLogLog)


## Using
```go

// Create HyperBitBit
hbb := hyperbitbit.NewHyperBitBit(6)

// Add value to HyperBitBit
hbb.Add([]byte("hello"))


// Returns cardinality
hbb.Get()

```
