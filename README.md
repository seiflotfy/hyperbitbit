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

## Initial Results
From [demo](demo)
```2017/02/20 12:44:46 
        file:  data/words-1 
        exact: 150 
        estimate: 1473 
        ratio: 89.816701%
2017/02/20 12:44:46 
        file:  data/words-2 
        exact: 1308 
        estimate: 2039 
        ratio: 35.850907%
2017/02/20 12:44:46 
        file:  data/words-3 
        exact: 76205 
        estimate: 68141 
        ratio: -11.834285%
2017/02/20 12:44:46 
        file:  data/words-4 
        exact: 235886 
        estimate: 244589 
        ratio: 3.558214%
2017/02/20 12:44:47 
        file:  data/words-5 
        exact: 349900 
        estimate: 317192 
        ratio: -10.311735%
2017/02/20 12:44:47 
        file:  data/words-6 
        exact: 479829 
        estimate: 510835 
        ratio: 6.069670%
THIS IS WAAAAAY OFF, DUE TO REINSERTIONS --> Needs FIXING
2017/02/20 12:44:47 
        total
        exact: 660131 
        estimate: 1114140 
        ratio: 40.749726%
 ```
