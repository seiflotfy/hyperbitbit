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
hbb := hyperbitbit.New()

// Add value to HyperBitBit
hbb.Add([]byte("hello"))

// Returns cardinality
hbb.Cardinality()
```

## Initial Results
From [demo](hbbdemo/main.go)
```
2017/03/21 23:16:49
        file:  data/words-1
        exact: 150
        estimate: 1380
        ratio: 89.13%
2017/03/21 23:16:49
        file:  data/words-2
        exact: 1308
        estimate: 1869
        ratio: 30.02%
2017/03/21 23:16:49
        file:  data/words-3
        exact: 76205
        estimate: 62486
        ratio: -21.96%
2017/03/21 23:16:49
        file:  data/words-4
        exact: 235886
        estimate: 255417
        ratio: 7.65%
2017/03/21 23:16:50
        file:  data/words-5
        exact: 349900
        estimate: 317192
        ratio: -10.31%
2017/03/21 23:16:50
        file:  data/words-6
        exact: 479829
        estimate: 448578
        ratio: -6.97%
2017/03/21 23:16:50
        total
        exact: 660131
        estimate: 648276
        ratio: -1.83%
 ```
